package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/pkg/httputils"
	"github.com/zerodoctor/shawarma/pkg/model"
	gdb "github.com/zerodoctor/shawarma/pkg/plugin/github/db"
	gmodel "github.com/zerodoctor/shawarma/pkg/plugin/github/model"
)

const (
	GITHUB_OAUTH_TOKEN_URL string = "https://github.com/login/oauth/access_token"

	GITHUB_API_ENDPOINT  string = "https://api.github.com"
	GITHUB_AUTH_USER_URL string = GITHUB_API_ENDPOINT + "/user"
)

var (
	ErrMissingGithubCode error = errors.New("bad format") // NOTE: we want to hide reason from user
	ErrFormatGithubCode  error = errors.New("bad format") // NOTE: we want to hide reason from user
)

type GithubService struct {
	db *gdb.DB
}

func NewService() *GithubService {
	return &GithubService{}
}

func (s *GithubService) Setup(db db.DB) {
	s.db = gdb.NewDB(db.GetType(), db.GetConnection())
}

func (s *GithubService) RegisterUser(details map[string]interface{}) (model.User, error) {
	var user model.User

	codeInter, ok := details["code"]
	if !ok {
		return user, ErrMissingGithubCode
	}

	code, ok := codeInter.(string)
	if !ok {
		return user, ErrFormatGithubCode
	}

	githubUser, err := s.SaveGithubAuthUser(code)
	if err != nil {
		return user, err
	}
	user.Name = githubUser.Name
	user.AvatarURL = githubUser.AvatarURL

	return user, nil
}

func (s *GithubService) SaveGithubAuthUser(code string) (gmodel.GithubUser, error) {
	var githubUser gmodel.GithubUser

	token, err := s.GetGithubToken(code)
	if err != nil {
		return githubUser, fmt.Errorf("failed to fetch github token [error=%w]", err)
	}

	githubUser, err = s.GetGithubAuthUser(token)
	if err != nil {
		return githubUser, fmt.Errorf("failed to fetch github user [error=%w]", err)
	}
	githubUser.Token = token

	githubOrgs, err := s.SaveGithubUserOrgs(token, githubUser)
	if err != nil {
		return githubUser, fmt.Errorf("failed to save github user orgs [error=%w]", err)
	}
	githubUser.Orgs = githubOrgs

	repos, err := s.SaveGithubUserRepos(token, githubUser)
	if err != nil {
		return githubUser, fmt.Errorf("failed to save github user repos [error=%w]", err)
	}
	githubUser.Repos = repos

	return githubUser, nil
}

func (s *GithubService) GetGithubToken(code string) (string, error) {
	url := fmt.Sprintf("%s?code=%s&client_id=%s&client_secret=%s",
		GITHUB_OAUTH_TOKEN_URL, code,
		os.Getenv("GITHUB_CLIENT_ID"), os.Getenv("GITHUB_CLIENT_SECRET"),
	)
	resp, err := NewRequest(httputils.POST, url, nil).Do()
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

	resp, err := NewRequest(httputils.GET, GITHUB_AUTH_USER_URL, nil).
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

func (s *GithubService) RegisterUserOrganizations(token string, user model.User) ([]model.Organization, error) {
	var orgs []model.Organization
	return orgs, nil
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
		org, err = s.db.SaveGithubUserOrg(user.ID, org)
		if err != nil {
			return githubOrgs, fmt.Errorf("failed to save [org=%s] [error=%w]", org.Name, err)
		}
		githubOrgs = append(githubOrgs, org)
	}

	return githubOrgs, nil
}

func (s *GithubService) GetGithubUserOrgs(token string, userOrgsURL string) ([]gmodel.GithubUserOrg, error) {
	var userOrgs []gmodel.GithubUserOrg

	resps, err := NewRequest(httputils.GET, userOrgsURL, nil).
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

	resp, err := NewRequest(httputils.GET, orgURL, nil).
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

func (s *GithubService) RegisterUserRepositories(token string, user model.User) ([]model.Repository, error) {
	var repos []model.Repository
	return repos, nil
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
		repo, err := s.db.SaveGithubRepo(repos[i])
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

	resps, err := NewRequest(httputils.GET, reposURL, nil).
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
