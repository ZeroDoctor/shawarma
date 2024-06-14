package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/zerodoctor/shawarma/internal/model"
)

const OAUTH_GITHUB_TOKEN_URL string = "https://github.com/login/oauth/access_token"

func GetGithubToken(user model.GithubUser) (string, error) {

	authURL, err := url.Parse(OAUTH_GITHUB_TOKEN_URL)
	if err != nil {
		return "", fmt.Errorf("failed to parse [url=%s] [error=%w]",
			OAUTH_GITHUB_TOKEN_URL, err,
		)
	}

	req, err := http.NewRequest("POST", authURL.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request [error=%w]", err)
	}
	req.Header.Add("Accept", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to do request [error=%w]", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read all response from [url=%s] [error=%w]",
			OAUTH_GITHUB_TOKEN_URL, err,
		)
	}

	var tokenResponse model.GithubTokenResponse
	if err := json.Unmarshal(data, &tokenResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal token response from [url=%s] [error=%w]",
			OAUTH_GITHUB_TOKEN_URL, err,
		)
	}

	log.Debugf("for [state=%s] receive token with [scope=%s]",
		user.State, tokenResponse.Scope,
	)
	return tokenResponse.AccessToken, nil
}
