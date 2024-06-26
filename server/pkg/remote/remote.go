package remote

import (
	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/pkg/model"
)

type GitRemote interface {
	RegisterUser(map[string]interface{}) (model.User, error)
	Setup(db db.DB)
}
