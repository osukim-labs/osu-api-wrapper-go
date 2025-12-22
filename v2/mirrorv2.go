package v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/osukim-labs/osu-api-wrapper-go/client"
	"github.com/osukim-labs/osu-api-wrapper-go/v2/interfaces"
)

var (
	ErrClientNotInitialized = errors.New("client not initialized")
	ErrInvalidProvider      = errors.New("invalid mirror provider")
)

// MirrorProvider represents available mirror providers
type MirrorProvider string

const (
	ProviderMino      MirrorProvider = "Mino"
	ProviderOsuDirect MirrorProvider = "osu!Direct"
)

type OsuV2Mirror struct {
	*Base

	MirrorProvider string
	MirrorHost     string
	MirrorAPIPath  string
	client         *client.Client
	mu             sync.RWMutex
	initialized    bool
}

// NewOsuV2Mirror creates and initializes a new OsuV2Mirror instance
func NewOsuV2Mirror(provider string, timeout time.Duration) (*OsuV2Mirror, error) {
	api := &OsuV2Mirror{
		MirrorProvider: provider,
		Base: &Base{
			Timeout: timeout,
		},
	}

	// Map the provider to host and path
	if err := api.MapMirrorProvider(); err != nil {
		return nil, err
	}

	// Initialize the client
	api.client = client.NewClient(api.MirrorHost+api.MirrorAPIPath, api.Timeout, "")
	api.initialized = true

	return api, nil
}

// NewOsuV2MirrorWithCustomHost creates a mirror instance with custom host and path
func NewOsuV2MirrorWithCustomHost(host, apiPath string, timeout time.Duration) *OsuV2Mirror {
	api := &OsuV2Mirror{
		MirrorHost:    host,
		MirrorAPIPath: apiPath,
		Base: &Base{
			Timeout: timeout,
		},
	}

	api.client = client.NewClient(api.MirrorHost+api.MirrorAPIPath, api.Timeout, "")
	api.initialized = true

	return api
}

// MapMirrorProvider maps provider name to host and API path
func (api *OsuV2Mirror) MapMirrorProvider() error {
	switch MirrorProvider(api.MirrorProvider) {
	case ProviderMino:
		api.MirrorHost = "https://catboy.best"
		api.MirrorAPIPath = "/api/v2"
	case ProviderOsuDirect:
		api.MirrorHost = "https://osu.direct"
		api.MirrorAPIPath = "/api/v2"
	default:
		return fmt.Errorf("%w: %s (available: Mino, osu!Direct)", ErrInvalidProvider, api.MirrorProvider)
	}
	return nil
}

// ensureInitialized checks if the client is initialized and initializes if needed
func (api *OsuV2Mirror) ensureInitialized() error {
	api.mu.RLock()
	if api.initialized && api.client != nil {
		api.mu.RUnlock()
		return nil
	}
	api.mu.RUnlock()

	// Need to initialize
	api.mu.Lock()
	defer api.mu.Unlock()

	// Double-check after acquiring write lock
	if api.initialized && api.client != nil {
		return nil
	}

	// Map provider if not done yet
	if api.MirrorHost == "" || api.MirrorAPIPath == "" {
		if err := api.MapMirrorProvider(); err != nil {
			return err
		}
	}

	// Initialize client
	if api.client == nil {
		if api.Base == nil || api.Timeout == 0 {
			return errors.New("timeout not configured")
		}
		api.client = client.NewClient(api.MirrorHost+api.MirrorAPIPath, api.Timeout, "")
	}

	api.initialized = true
	return nil
}

// GetBeatmap retrieves a beatmap by ID
func (api *OsuV2Mirror) GetBeatmap(beatmapId string) (interfaces.Beatmap, error) {
	return api.GetBeatmapWithContext(context.Background(), beatmapId)
}

// GetBeatmapSet retrieves a beatmap set by ID
func (api *OsuV2Mirror) GetBeatmapSet(beatmapSetId string) (interfaces.BeatmapSet, error) {
	return api.GetBeatmapSetWithContext(context.Background(), beatmapSetId)
}

// GetBeatmapWithContext retrieves a beatmap by ID with context support
func (api *OsuV2Mirror) GetBeatmapWithContext(ctx context.Context, beatmapId string) (interfaces.Beatmap, error) {
	var beatmap interfaces.Beatmap

	// Validate input
	if err := validateBeatmapID(beatmapId); err != nil {
		return beatmap, err
	}

	// Ensure client is initialized
	if err := api.ensureInitialized(); err != nil {
		return beatmap, fmt.Errorf("initialization failed: %w", err)
	}

	// Get client safely
	api.mu.RLock()
	req := api.client
	api.mu.RUnlock()

	if req == nil {
		return beatmap, ErrClientNotInitialized
	}

	// Make request
	rawBody, statusCode, err := req.SendGetRequestWithContext(ctx, "/b/"+beatmapId)
	if err != nil {
		return beatmap, fmt.Errorf("request failed: %w", err)
	}

	// Check status code
	if statusCode != 200 {
		return beatmap, req.HandleError(statusCode, string(rawBody))
	}

	// Parse response
	if err := json.Unmarshal(rawBody, &beatmap); err != nil {
		return beatmap, fmt.Errorf("failed to parse beatmap: %w", err)
	}

	return beatmap, nil
}

// GetBeatmapSetWithContext retrieves a beatmap set by ID with context support
func (api *OsuV2Mirror) GetBeatmapSetWithContext(ctx context.Context, beatmapSetId string) (interfaces.BeatmapSet, error) {
	var beatmapset interfaces.BeatmapSet

	// Validate input
	if err := validateBeatmapID(beatmapSetId); err != nil {
		return beatmapset, err
	}

	// Ensure client is initialized
	if err := api.ensureInitialized(); err != nil {
		return beatmapset, fmt.Errorf("initialization failed: %w", err)
	}

	// Get client safely
	api.mu.RLock()
	req := api.client
	api.mu.RUnlock()

	if req == nil {
		return beatmapset, ErrClientNotInitialized
	}

	// Make request
	rawBody, statusCode, err := req.SendGetRequestWithContext(ctx, "/s/"+beatmapSetId)
	if err != nil {
		return beatmapset, fmt.Errorf("request failed: %w", err)
	}

	// Check status code
	if statusCode != 200 {
		return beatmapset, req.HandleError(statusCode, string(rawBody))
	}

	// Parse response
	if err := json.Unmarshal(rawBody, &beatmapset); err != nil {
		return beatmapset, fmt.Errorf("failed to parse beatmapset: %w", err)
	}

	return beatmapset, nil
}

// IsInitialized returns whether the client has been initialized
func (api *OsuV2Mirror) IsInitialized() bool {
	api.mu.RLock()
	defer api.mu.RUnlock()
	return api.initialized && api.client != nil
}

// GetProviderInfo returns information about the current mirror provider
func (api *OsuV2Mirror) GetProviderInfo() (provider, host, path string) {
	api.mu.RLock()
	defer api.mu.RUnlock()
	return api.MirrorProvider, api.MirrorHost, api.MirrorAPIPath
}
