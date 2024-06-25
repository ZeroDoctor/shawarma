package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"

	"github.com/zerodoctor/shawarma/pkg/model"
	"github.com/zerodoctor/shawarma/pkg/plugin/github/db"
	gmodel "github.com/zerodoctor/shawarma/pkg/plugin/github/model"
	"github.com/zerodoctor/shawarma/pkg/service"
)

const (
	GITHUB_OAUTH_TOKEN_URL string = "https://github.com/login/oauth/access_token"

	GITHUB_API_ENDPOINT  string = "https://api.github.com"
	GITHUB_AUTH_USER_URL string = GITHUB_API_ENDPOINT + "/user"
)

type GithubService struct {
	db db.DB
}

func NewService(db db.DB) *GithubService {
	return &GithubService{
		db: db,
	}
}

func (s *GithubService) SaveGithubAuthUser(secret string, user model.User) (model.User, error) {
	token, err := s.GetGithubToken(secret, user)
	if err != nil {
		return user, fmt.Errorf("failed to fetch github token [error=%w]", err)
	}

	githubUser, err := s.GetGithubAuthUser(token)
	if err != nil {
		return user, fmt.Errorf("failed to fetch github user [error=%w]", err)
	}

	githubOrgs, err := s.SaveGithubUserOrgs(token, githubUser)
	if err != nil {
		return user, fmt.Errorf("failed to save github user orgs [error=%w]", err)
	}
	githubUser.Orgs = githubOrgs

	repos, err := s.SaveGithubUserRepos(token, githubUser)
	if err != nil {
		return user, fmt.Errorf("failed to save github user repos [error=%w]", err)
	}
	githubUser.Repos = repos

	return user, nil
}

func (s *GithubService) GetGithubToken(secret string, user model.User) (string, error) {
	resp, err := NewRequest(service.HTTP_POST, GITHUB_OAUTH_TOKEN_URL, nil).
		OptionGithubHeaders(secret).
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

	var tokenResponse gmodel.GithubTokenResponse
	if err := json.Unmarshal(data, &tokenResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal token response from [url=%s] [error=%w]",
			GITHUB_OAUTH_TOKEN_URL, err,
		)
	}

	return tokenResponse.AccessToken, nil
}

func (s *GithubService) GetGithubAuthUser(token string) (gmodel.GithubUser, error) {
	var user gmodel.GithubUser

	resp, err := NewRequest(service.HTTP_GET, GITHUB_AUTH_USER_URL, nil).
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

func (s *GithubService) SaveGithubUserOrgs(token string, user gmodel.GithubUser) ([]gmodel.GithubOrg, error) {
	var githubOrgs []gmodel.GithubOrg

	githubUserOrgs, err := s.GetGithubUserOrgs(token, user.OrganizationsURL)
	if err != nil {
		return githubOrgs, fmt.Errorf("failed to fetch github user orgs [error=%w]", err)
	}

	for i := range githubUserOrgs {
		org, err := s.GetGithubOrg(token, githubUserOrgs[i].URL)
		if err != nil {
			return githubOrgs, fmt.Errorf("failed to fetch github [org=%s] [error=%w]",
				githubUserOrgs[i].Login, err,
			)
		}
		org, err = s.db.InsertGithubUserOrg(user.ID, org)
		if err != nil {
			return githubOrgs, fmt.Errorf("failed to save [org=%s] [error=%w]", org.Name, err)
		}
		githubOrgs = append(githubOrgs, org)
	}

	return githubOrgs, nil
}

func (s *GithubService) GetGithubUserOrgs(token string, userOrgsURL string) ([]gmodel.GithubUserOrg, error) {
	var userOrgs []gmodel.GithubUserOrg

	resps, err := NewRequest(service.HTTP_GET, userOrgsURL, nil).
		OptionGithubHeaders(token).
		OptionGithubPages(100).
		DoAll()
	if err != nil {
		return userOrgs, fmt.Errorf("failed to do request [error=%w]", err)
	}

	for _, resp := range resps {
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return userOrgs, fmt.Errorf("failed to read all response from [url=%s] [error=%w]",
				userOrgsURL, err,
			)
		}

		var orgs []gmodel.GithubUserOrg
		if err := json.Unmarshal(data, &orgs); err != nil {
			return userOrgs, fmt.Errorf("failed to unmarshal data from github user orgs [error=%w]", err)
		}
		userOrgs = append(userOrgs, orgs...)
	}

	return userOrgs, nil
}

func (s *GithubService) GetGithubOrg(token string, orgURL string) (gmodel.GithubOrg, error) {
	var org gmodel.GithubOrg

	resp, err := NewRequest(service.HTTP_GET, orgURL, nil).
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

	sql.Drivers()

	if err := json.Unmarshal(data, &org); err != nil {
		return org, fmt.Errorf("failed to unmarshal data from github org [error=%w]", err)
	}

	return org, nil
}

func (s *GithubService) SaveGithubUserRepos(token string, user gmodel.GithubUser) ([]gmodel.GithubRepo, error) {
	var repos []gmodel.GithubRepo

	userRepos, err := s.GetGithubRepos(token, user.ReposURL)
	if err != nil {
		return repos, fmt.Errorf("failed to fetch [user=%s] repos [error=%w]", user.Name, err)
	}
	repos = append(repos, userRepos...)

	for i := range user.Orgs {
		orgRepos, err := s.GetGithubRepos(token, user.Orgs[i].ReposURL)
		if err != nil {
			return repos, fmt.Errorf("failed to fetch [user=%s] [org=%s] repos [error=%w]",
				user.Name, user.Orgs[i].Name, err,
			)
		}
		repos = append(repos, orgRepos...)
	}

	for i := range repos {
		repo, err := s.db.InsertGithubRepo(repos[i])
		if err != nil {
			return repos, fmt.Errorf("failed to save [user=%s] [owner=%s] repo [error=%w]",
				user.Name, repos[i].Owner.Login, err,
			)
		}
		repos[i] = repo
	}

	return repos, nil
}

func (s *GithubService) GetGithubRepos(token string, reposURL string) ([]gmodel.GithubRepo, error) {
	var repos []gmodel.GithubRepo

	resps, err := NewRequest(service.HTTP_GET, reposURL, nil).
		OptionGithubHeaders(token).
		OptionGithubPages(100).
		DoAll()
	if err != nil {
		return repos, fmt.Errorf("failed to do request [error=%w]", err)
	}
	for _, resp := range resps {
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return repos, fmt.Errorf("failed to read all response from [url=%s] [error=%w]",
				reposURL, err,
			)
		}

		var tempRepos []gmodel.GithubRepo
		if err := json.Unmarshal(data, &tempRepos); err != nil {
			return repos, fmt.Errorf("failed to unmarshal data from github repos [error=%w]", err)
		}
		repos = append(repos, tempRepos...)
	}

	return repos, nil
}
