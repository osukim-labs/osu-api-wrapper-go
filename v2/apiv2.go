package v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/osukim-labs/osu-api-wrapper-go/client"
	"github.com/osukim-labs/osu-api-wrapper-go/v2/interfaces"
)

var (
	ErrTokenExpired       = errors.New("access token expired")
	ErrInvalidCredentials = errors.New("invalid oauth credentials")
	ErrEmptyBeatmapList   = errors.New("beatmap list is empty")
)

type OsuV2API struct {
	*Base
	OauthBody
	OauthResp OauthResp
	client    *client.Client
	mu        sync.RWMutex
}

// NewOsuV2API creates a new OsuV2API instance with the given credentials
func NewOsuV2API(clientID int, clientSecret string, timeout time.Duration) (*OsuV2API, error) {
	if clientID == 0 {
		return nil, fmt.Errorf("%w: client_id is required", ErrInvalidCredentials)
	}
	if clientSecret == "" {
		return nil, fmt.Errorf("%w: client_secret is required", ErrInvalidCredentials)
	}

	api := &OsuV2API{
		Base: &Base{
			Timeout: timeout,
		},
		OauthBody: OauthBody{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			GrantType:    "client_credentials",
			Scope:        "public",
		},
	}

	// Perform initial login
	if err := api.Login(); err != nil {
		return nil, fmt.Errorf("initial login failed: %w", err)
	}

	return api, nil
}

// NewOsuV2APIWithCustomScope creates a new OsuV2API instance with custom scope
func NewOsuV2APIWithCustomScope(clientID int, clientSecret, scope string, timeout time.Duration) (*OsuV2API, error) {
	if clientID == 0 {
		return nil, fmt.Errorf("%w: client_id is required", ErrInvalidCredentials)
	}
	if clientSecret == "" {
		return nil, fmt.Errorf("%w: client_secret is required", ErrInvalidCredentials)
	}

	api := &OsuV2API{
		Base: &Base{
			Timeout: timeout,
		},
		OauthBody: OauthBody{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			GrantType:    "client_credentials",
			Scope:        scope,
		},
	}

	// Perform initial login
	if err := api.Login(); err != nil {
		return nil, fmt.Errorf("initial login failed: %w", err)
	}

	return api, nil
}

// IsTokenExpired checks if the current access token has expired
func (api *OsuV2API) IsTokenExpired() bool {
	api.mu.RLock()
	defer api.mu.RUnlock()

	if api.OauthResp.CreatedAt.IsZero() {
		return true
	}

	// Add 5-minute buffer before actual expiration
	expiresAt := api.OauthResp.CreatedAt.Add(time.Duration(api.OauthResp.ExpiresIn-300) * time.Second)
	return time.Now().After(expiresAt)
}

// GetTokenExpiresIn returns the duration until the token expires
func (api *OsuV2API) GetTokenExpiresIn() time.Duration {
	api.mu.RLock()
	defer api.mu.RUnlock()

	if api.OauthResp.CreatedAt.IsZero() {
		return 0
	}

	expiresAt := api.OauthResp.CreatedAt.Add(time.Duration(api.OauthResp.ExpiresIn) * time.Second)
	remaining := time.Until(expiresAt)

	if remaining < 0 {
		return 0
	}
	return remaining
}

// IsTokenValid checks if the token is valid and refreshes it if needed
func (api *OsuV2API) IsTokenValid() error {
	if !api.IsTokenExpired() {
		return nil
	}

	// Double-check locking pattern for thread safety
	api.mu.Lock()
	defer api.mu.Unlock()

	// Check again after acquiring lock
	if !api.IsTokenExpired() {
		return nil
	}

	return api.login()
}

// Login performs OAuth authentication and obtains an access token
func (api *OsuV2API) Login() error {
	api.mu.Lock()
	defer api.mu.Unlock()
	return api.login()
}

// login is the internal login method (must be called with lock held)
func (api *OsuV2API) login() error {
	req := client.NewClient(OauthUrl, api.Timeout, "")

	data := url.Values{}
	data.Set("client_id", strconv.Itoa(api.ClientID))
	data.Set("client_secret", api.ClientSecret)
	data.Set("grant_type", api.GrantType)
	data.Set("scope", api.Scope)

	rawBody, statusCode, err := req.OauthRequest(data)
	if err != nil {
		return fmt.Errorf("oauth request failed: %w", err)
	}

	if statusCode != 200 {
		return fmt.Errorf("oauth failed (status %d): %w", statusCode, req.HandleError(statusCode, string(rawBody)))
	}

	if err := json.Unmarshal(rawBody, &api.OauthResp); err != nil {
		return fmt.Errorf("failed to parse oauth response: %w", err)
	}

	api.OauthResp.CreatedAt = time.Now()
	api.client = client.NewClient(APIUrl, api.Timeout, api.OauthResp.AccessToken)

	return nil
}

// GetBeatmap retrieves a beatmap by ID
func (api *OsuV2API) GetBeatmap(beatmapId string) (interfaces.Beatmap, error) {
	return api.GetBeatmapWithContext(context.Background(), beatmapId)
}

// GetBeatmapWithContext retrieves a beatmap by ID with context support
func (api *OsuV2API) GetBeatmapWithContext(ctx context.Context, beatmapId string) (interfaces.Beatmap, error) {
	var beatmap interfaces.Beatmap

	if err := validateBeatmapID(beatmapId); err != nil {
		return beatmap, err
	}

	if err := api.IsTokenValid(); err != nil {
		return beatmap, fmt.Errorf("token validation failed: %w", err)
	}

	api.mu.RLock()
	req := api.client
	api.mu.RUnlock()

	if req == nil {
		return beatmap, errors.New("client not initialized")
	}

	rawBody, statusCode, err := req.SendGetRequestWithContext(ctx, "/beatmaps/"+beatmapId)
	if err != nil {
		return beatmap, fmt.Errorf("request failed: %w", err)
	}

	if statusCode != 200 {
		return beatmap, req.HandleError(statusCode, string(rawBody))
	}

	if err := json.Unmarshal(rawBody, &beatmap); err != nil {
		return beatmap, fmt.Errorf("failed to parse beatmap: %w", err)
	}

	return beatmap, nil
}

// GetBeatmapSet retrieves a beatmap set by ID
func (api *OsuV2API) GetBeatmapSet(beatmapSetId string) (interfaces.BeatmapSet, error) {
	return api.GetBeatmapSetWithContext(context.Background(), beatmapSetId)
}

// GetBeatmapSetWithContext retrieves a beatmap set by ID with context support
func (api *OsuV2API) GetBeatmapSetWithContext(ctx context.Context, beatmapSetId string) (interfaces.BeatmapSet, error) {
	var beatmapset interfaces.BeatmapSet

	if err := validateBeatmapID(beatmapSetId); err != nil {
		return beatmapset, err
	}

	if err := api.IsTokenValid(); err != nil {
		return beatmapset, fmt.Errorf("token validation failed: %w", err)
	}

	api.mu.RLock()
	req := api.client
	api.mu.RUnlock()

	if req == nil {
		return beatmapset, errors.New("client not initialized")
	}

	rawBody, statusCode, err := req.SendGetRequestWithContext(ctx, "/beatmapsets/"+beatmapSetId)
	if err != nil {
		return beatmapset, fmt.Errorf("request failed: %w", err)
	}

	if statusCode != 200 {
		return beatmapset, req.HandleError(statusCode, string(rawBody))
	}

	if err := json.Unmarshal(rawBody, &beatmapset); err != nil {
		return beatmapset, fmt.Errorf("failed to parse beatmapset: %w", err)
	}

	return beatmapset, nil
}

// GetBeatmaps retrieves multiple beatmaps by their IDs
func (api *OsuV2API) GetBeatmaps(beatmapIds []string) ([]interfaces.Beatmap, error) {
	return api.GetBeatmapsWithContext(context.Background(), beatmapIds)
}

// GetBeatmapsWithContext retrieves multiple beatmaps by their IDs with context support
func (api *OsuV2API) GetBeatmapsWithContext(ctx context.Context, beatmapIds []string) ([]interfaces.Beatmap, error) {
	if len(beatmapIds) == 0 {
		return nil, ErrEmptyBeatmapList
	}

	var beatmapResp interfaces.BeatmapsResp

	if err := api.IsTokenValid(); err != nil {
		return nil, fmt.Errorf("token validation failed: %w", err)
	}

	api.mu.RLock()
	req := api.client
	api.mu.RUnlock()

	if req == nil {
		return nil, errors.New("client not initialized")
	}

	// Build query string with validated IDs only
	var sb strings.Builder
	validCount := 0
	for _, id := range beatmapIds {
		if err := validateBeatmapID(id); err != nil {
			continue
		}

		if validCount > 0 {
			sb.WriteString("&")
		}
		sb.WriteString("ids[]=")
		sb.WriteString(id)
		validCount++
	}

	if validCount == 0 {
		return nil, errors.New("no valid beatmap IDs provided")
	}

	rawBody, statusCode, err := req.SendGetRequestWithContext(ctx, "/beatmaps?"+sb.String())
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if statusCode != 200 {
		return nil, req.HandleError(statusCode, string(rawBody))
	}

	if err := json.Unmarshal(rawBody, &beatmapResp); err != nil {
		return nil, fmt.Errorf("failed to parse beatmaps: %w", err)
	}

	return beatmapResp.Beatmaps, nil
}

// IsAuthenticated returns whether the API client has a valid token
func (api *OsuV2API) IsAuthenticated() bool {
	api.mu.RLock()
	defer api.mu.RUnlock()
	return api.client != nil && !api.IsTokenExpired()
}

// GetAccessToken returns the current access token (useful for debugging)
func (api *OsuV2API) GetAccessToken() string {
	api.mu.RLock()
	defer api.mu.RUnlock()
	return api.OauthResp.AccessToken
}

// RefreshToken manually refreshes the OAuth token
func (api *OsuV2API) RefreshToken() error {
	return api.Login()
}
