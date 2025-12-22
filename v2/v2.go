package v2

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const BaseUrl string = "https://osu.ppy.sh"
const APIUrl string = BaseUrl + "/api/v2"
const OauthUrl string = BaseUrl + "/oauth/token"

type Base struct {
	Timeout time.Duration
}

type OauthBody struct {
	ClientID     int    `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type OauthResp struct {
	CreatedAt    time.Time
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
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
