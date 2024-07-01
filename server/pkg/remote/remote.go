package remote

import (
	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/pkg/model"
)

type GitRemoteDriver interface {
	Setup(db.DB)

	RegisterUser(map[string]interface{}) (model.User, error)
	RegisterUserOrganizations(model.User) ([]model.Organization, error)
	RegisterUserRepositories(model.User) ([]model.Repository, error)
	GetCommitsURL(model.User, []string) ([]model.Commit, error)
}
