package remote

import (
	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/pkg/model"
)

type GitRemoteDriver interface {
	Setup(db.DB)

	RegisterUser(token string) (model.User, error)
	RegisterUserOrganizations(string, model.User) ([]model.Organization, error)
	RegisterUserRepositories(string, model.User) ([]model.Repository, error)

	FetchNewToken(model.UserGitRegisterDetails) (string, error)

	GetCommitsURL(string, model.User, []string) ([]model.Commit, error)
}
