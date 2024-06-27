package remote

import (
	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/pkg/model"
)

type GitRemote interface {
	Setup(db db.DB)

	RegisterUser(map[string]interface{}) (model.User, error)
	RegisterUserOrganizations(token string, user model.User) ([]model.Organization, error)
	RegisterUserRepositories(token string, user model.User) ([]model.Repository, error)
}
