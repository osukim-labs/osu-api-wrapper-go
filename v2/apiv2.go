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

type OsuV2API struct {
	*Base
	OauthBody
	OauthResp OauthResp
	client    *client.Client
	mu        sync.RWMutex
}

func (api *OsuV2API) IsTokenExpired() bool {
	api.mu.RLock()
	defer api.mu.RUnlock()

	if api.OauthResp.CreatedAt.IsZero() {
		return true
	}

	expiresAt := api.OauthResp.CreatedAt.Add(time.Duration(api.OauthResp.ExpiresIn-300) * time.Second)
	return time.Now().After(expiresAt)
}

func (api *OsuV2API) IsTokenValid() error {
	if !api.IsTokenExpired() {
		return nil
	}

	// Check if token expired after getting the lock
	api.mu.Lock()
	defer api.mu.Unlock()

	if !api.IsTokenExpired() {
		return nil
	}

	return api.login()
}

func (api *OsuV2API) Login() error {
	api.mu.Lock()
	defer api.mu.Unlock()
	return api.login()
}

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
		return req.HandleError(statusCode, string(rawBody))
	}

	if err := json.Unmarshal(rawBody, &api.OauthResp); err != nil {
		return fmt.Errorf("failed to parse oauth response: %w", err)
	}

	api.OauthResp.CreatedAt = time.Now()

	api.client = client.NewClient(APIUrl, api.Timeout, api.OauthResp.AccessToken)

	return nil
}

func (api *OsuV2API) GetBeatmap(beatmapId string) (interfaces.Beatmap, error) {
	return api.GetBeatmapWithContext(context.Background(), beatmapId)
}

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

func (api *OsuV2API) GetBeatmapSet(beatmapSetId string) (interfaces.BeatmapSet, error) {
	return api.GetBeatmapSetWithContext(context.Background(), beatmapSetId)
}

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

func (api *OsuV2API) GetBeatmaps(beatmapIds []string) ([]interfaces.Beatmap, error) {
	return api.GetBeatmapsWithContext(context.Background(), beatmapIds)
}

func (api *OsuV2API) GetBeatmapsWithContext(ctx context.Context, beatmapIds []string) ([]interfaces.Beatmap, error) {
	var beatmapResp interfaces.BeatmapsResp

	if err := api.IsTokenValid(); err != nil {
		return nil, fmt.Errorf("token validation failed: %w", err)
	}

	api.mu.RLock()
	req := api.client
	api.mu.RUnlock()

	var sb strings.Builder
	for i, id := range beatmapIds {
		if err := validateBeatmapID(id); err != nil {
			continue
		}

		if i > 0 {
			sb.WriteString("&")
		}
		sb.WriteString("ids[]=")
		sb.WriteString(id)
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

func validateBeatmapID(id string) error {
	if id == "" {
		return errors.New("beatmap ID cannot be empty")
	}
	if _, err := strconv.Atoi(id); err != nil {
		return fmt.Errorf("invalid beatmap ID: %w", err)
	}
	return nil
}
