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

func (s *Service) SaveGithubAuthUser(user model.User) (model.User, error) {
	token, err := s.GetGithubToken(user)
	if err != nil {
		return user, fmt.Errorf("failed to fetch github token [error=%w]", err)
	}
	user.GithubToken = token

	githubUser, err := s.GetGithubAuthUser(token)
	if err != nil {
		return user, fmt.Errorf("failed to fetch github user [error=%w]", err)
	}
	user.GithubUser = githubUser

	githubUserOrgs, err := s.GetGithubUserOrgs(token, githubUser.OrganizationsURL)
	if err != nil {
		return user, fmt.Errorf("failed to fetch github user orgs [error=%w]", err)
	}

	var githubOrgs []model.GithubOrg
	for i := range githubUserOrgs {
		org, err := s.GetGithubOrg(token, githubUserOrgs[i].URL)
		if err != nil {
			return user, fmt.Errorf("failed to fetch github [org=%s] [error=%w]",
				githubUserOrgs[i].Login, err,
			)
		}
		githubOrgs = append(githubOrgs, org)
	}
	// TODO: save githubOrgs with github user id
	//	also insert new row for org with generic values

	user, err = s.db.InsertUser(user)
	if err != nil {
		return user, fmt.Errorf("failed to save user [error=%w]", err)
	}
	user.GithubToken = ""

	return user, nil
}

func (s *Service) GetGithubToken(user model.User) (string, error) {
	resp, err := NewRequest(HTTP_POST, GITHUB_OAUTH_TOKEN_URL, nil).
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

func (s *Service) GetGithubAuthUser(token string) (model.GithubUser, error) {
	var user model.GithubUser

	resp, err := NewRequest(HTTP_GET, GITHUB_AUTH_USER_URL, nil).
		OptionGithubHeaders(token).
		Do()
	if err != nil {
		return user, fmt.Errorf("failed to do request [error=%w]", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return user, fmt.Errorf("failed to read all response from [url=%s] [error=%w]",
			GITHUB_AUTH_USER_URL, err,
		)
	}

	if err := json.Unmarshal(data, &user); err != nil {
		return user, fmt.Errorf("failed to unmarshal data for github auth user [error=%w]", err)
	}

	return user, nil
}

func (s *Service) GetGithubUserOrgs(token string, userOrgsURL string) ([]model.GithubUserOrg, error) {
	var userOrgs []model.GithubUserOrg

	resp, err := NewRequest(HTTP_GET, userOrgsURL, nil).
		OptionGithubHeaders(token).
		Do()
	if err != nil {
		return userOrgs, fmt.Errorf("failed to do request [error=%w]", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return userOrgs, fmt.Errorf("failed to read all response from [url=%s] [error=%w]",
			userOrgsURL, err,
		)
	}

	if err := json.Unmarshal(data, &userOrgs); err != nil {
		return userOrgs, fmt.Errorf("failed to unmarshal data from github user orgs [error=%w]", err)
	}

	return userOrgs, nil
}

func (s *Service) GetGithubOrg(token string, orgURL string) (model.GithubOrg, error) {
	var org model.GithubOrg

	resp, err := NewRequest(HTTP_GET, orgURL, nil).
		OptionGithubHeaders(token).
		Do()
	if err != nil {
		return org, fmt.Errorf("failed to do request [error=%w]", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return org, fmt.Errorf("failed to read all response from [url=%s] [error=%w]",
			orgURL, err,
		)
	}

	if err := json.Unmarshal(data, &org); err != nil {
		return org, fmt.Errorf("failed to unmarshal data from github org [error=%w]", err)
	}

	return org, nil
}
