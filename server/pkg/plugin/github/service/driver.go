package service

import (
	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/pkg/model"
	gdb "github.com/zerodoctor/shawarma/pkg/plugin/github/db"
)

type GithubDriver struct {
	*GithubService
}

func NewGithubDriver() *GithubDriver {
	return &GithubDriver{
		GithubService: &GithubService{
			GithubRequestLimiter: NewRequestLimiter(),
		},
	}
}

func (gd *GithubDriver) Setup(db db.DB) {
	gd.db = gdb.NewDB(db.GetType(), db.GetConnection())
}

func (gd *GithubDriver) RegisterUser(details map[string]interface{}) (model.User, error) {
	var user model.User

	codeInter, ok := details["code"]
	if !ok {
		return user, ErrMissingGithubCode
	}

	code, ok := codeInter.(string)
	if !ok {
		return user, ErrFormatGithubCode
	}

	githubUser, err := gd.SaveGithubAuthUser(code)
	if err != nil {
		return user, err
	}
	user.Name = githubUser.Login
	user.AvatarURL = githubUser.AvatarURL

	return user, nil
}

func (gd *GithubDriver) RegisterUserOrganizations(user model.User) ([]model.Organization, error) {
	var orgs []model.Organization

	githubUsers, err := gd.db.GetGithubUserByName(user.Name)
	if err != nil {
		return orgs, err
	}
	githubUser := githubUsers[0]

	for i := range githubUser.Orgs {
		orgs = append(orgs, model.Organization{
			Name:      githubUser.Orgs[i].Login,
			AvatarURL: githubUser.Orgs[i].AvatarURL,
		})
	}

	return orgs, nil
}

func (gd *GithubDriver) RegisterUserRepositories(user model.User) ([]model.Repository, error) {
	var repos []model.Repository

	githubUsers, err := gd.db.GetGithubUserByName(user.Name)
	if err != nil {
		return repos, err
	}
	githubUser := githubUsers[0]

	for i := range githubUser.Repos {
		var branches []model.Branch
		githubBranches := githubUser.Repos[i].Branches
		for j := range githubBranches {
			branches = append(branches, model.Branch{
				Name: githubBranches[j].Name,
				Hash: githubBranches[j].SHA,
			})
		}

		repos = append(repos, model.Repository{
			Name:          githubUser.Repos[i].Name,
			Owner:         githubUser.Repos[i].Owner.Login,
			OwnerType:     githubUser.Repos[i].Owner.Type,
			DefaultBranch: githubUser.Repos[i].DefaultBranch,
			Branches:      branches,
		})
	}
	return repos, nil
}

func (gd *GithubDriver) GetCommitsURL(user model.User, hashes []string) ([]model.Commit, error) {
	var commits []model.Commit
	// TODO: implementation
	return commits, nil
}
