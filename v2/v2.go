package v2

import "time"

const BaseUrl string = "https://osu.ppy.sh"
const APIUrl string = BaseUrl + "/api/v2"
const OauthUrl string = BaseUrl + "/oauth/token"

type Base struct {
	Timeout int
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
