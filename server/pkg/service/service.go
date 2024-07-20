package service

import (
	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/internal/logger"
	"github.com/zerodoctor/shawarma/pkg/model"
	_ "github.com/zerodoctor/shawarma/pkg/plugin/github"
	"github.com/zerodoctor/shawarma/pkg/remote"
)

var log = logger.Log

type Service struct {
	db        db.DB
	userPolls map[string][]Poll
}

func NewService(db db.DB) *Service {
	remote.Setup(db)

	return &Service{
		db:        db,
		userPolls: make(map[string][]Poll),
	}
}

func (s *Service) RegisterUser(remoteName string, details map[string]interface{}) (model.User, error) {
	service := remote.GetRemoteService(remoteName)
	user, err := service.RegisterUser(details)
	if err != nil {
		log.Errorf("[user=%s] failed to register with [remote=%s]", user.Name, remoteName)
		return user, err
	}

	log.Infof("[user=%s] registered with [remote=%s]", user.Name, remoteName)

	return user, err
}

func (s *Service) GetUser(name string) (model.User, error) {
	return s.db.QueryUserByName(name)
}

func (s *Service) GetUserPolls(user model.User) []Poll {
	var polls []Poll

	return polls
}
