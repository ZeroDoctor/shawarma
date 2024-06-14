package model

type GithubUser struct {
	State    string `json:"state"`
	ClientID string `json:"client_id"`
	Code     string `json:"code"`
	Token    string `json:"token"`
}

type GithubTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}
