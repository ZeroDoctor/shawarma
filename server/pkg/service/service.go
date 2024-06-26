package service

import (
	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/pkg/model"
	_ "github.com/zerodoctor/shawarma/pkg/plugin/github"
	"github.com/zerodoctor/shawarma/pkg/remote"
)

type Service struct {
	db db.DB
}

func NewService(db db.DB) *Service {
	remote.Setup(db)

	return &Service{
		db: db,
	}
}

func (s *Service) RegisterUser(remoteName string, details map[string]interface{}) (model.User, error) {
	return remote.GetRemoteService(remoteName).RegisterUser(details)
}
