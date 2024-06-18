package service

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/zerodoctor/shawarma/internal/model"
)

const (
	GITHUB_OAUTH_TOKEN_URL string = "https://github.com/login/oauth/access_token"

	GITHUB_API_ENDPOINT  string = "https://api.github.com"
	GITHUB_AUTH_USER_URL string = GITHUB_API_ENDPOINT + "/user"
)

func GetGithubToken(user model.User) (string, error) {
	resp, err := NewRequest(HTTP_POST, GITHUB_OAUTH_TOKEN_URL).
		OptionGithubHeaders(user.GithubToken).
		Do()
	if err != nil {
		return "", fmt.Errorf("failed to do request [error=%w]", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read all response from [url=%s] [error=%w]",
			GITHUB_OAUTH_TOKEN_URL, err,
		)
	}

	var tokenResponse model.GithubTokenResponse
	if err := json.Unmarshal(data, &tokenResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal token response from [url=%s] [error=%w]",
			GITHUB_OAUTH_TOKEN_URL, err,
		)
	}

	log.Debugf("for [state=%s] receive token with [scope=%s]",
		user.GithubState, tokenResponse.Scope,
	)
	return tokenResponse.AccessToken, nil
}

func GetGithubAuthUser(token string, user model.GithubAuthUser) (model.GithubAuthUser, error) {
	resp, err := NewRequest(HTTP_GET, GITHUB_AUTH_USER_URL).
		OptionGithubHeaders(token).
		Do()
	if err != nil {
		return user, fmt.Errorf("failed to do request [error=%w]", err)
	}
	defer resp.Body.Close()

	// TODO: figure out how to store data
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return user, fmt.Errorf("failed to read all response from [url=%s] [error=%w]",
			GITHUB_AUTH_USER_URL, err,
		)
	}

	return user, nil
}
