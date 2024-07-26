package service

import (
	"context"
	"fmt"

	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/internal/logger"
	"github.com/zerodoctor/shawarma/pkg/model"
	_ "github.com/zerodoctor/shawarma/pkg/plugin/github"
	"github.com/zerodoctor/shawarma/pkg/remote"
)

var log = logger.Log

type Service struct {
	db        db.DB
	userPolls map[string][]*Poll
}

func NewService(db db.DB) *Service {
	remote.Setup(db)

	return &Service{
		db:        db,
		userPolls: make(map[string][]*Poll),
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

	if userCount, err := s.db.QueryUserCount(); userCount != 0 || err != nil {
		return user, err
	}
	user.IsOwner = true

	user, err = s.db.SaveUser(user)
	if err != nil {
		log.Errorf("[user=%s] failed to save user [error=%s]", user.Name, err.Error())
		return user, err
	}

	orgPoll := NewPoll(context.Background(), "/orgs", func(ctx context.Context, p *Poll) error {
		var errs []error

		orgs, err := service.RegisterUserOrganizations(user.Tokens[remoteName], user)
		if err != nil {
			log.Errorf("[user=%s] failed to register orgs [error=%s]", user.Name, err.Error())
			errs = append(errs, err)
		}

		for i := range orgs {
			_, err := s.db.SaveOrganization(orgs[i])
			if err != nil {
				log.Errorf("[user=%s] failed to save [org=%s] [error=%s]", user.Name, orgs[i].Name, err.Error())
				errs = append(errs, fmt.Errorf("failed save [org=%s] for [user=%s] [error=%w]",
					orgs[i].Name, user.Name, err,
				))
			}

			log.Infof("[user=%s] [orgs=%s] [remote=%s] saved", user.Name, orgs[i].Name, remoteName)
		}

		return combindErr(errs)
	})
	s.userPolls[user.Name] = append(s.userPolls[user.Name], orgPoll)

	repoPoll := NewPoll(context.Background(), "/repos", func(ctx context.Context, p *Poll) error {
		var errs []error

		repos, err := service.RegisterUserRepositories(user.Tokens[remoteName], user)
		if err != nil {
			log.Errorf("[user=%s] failed to register repos [error=%s]", user.Name, err.Error())
			errs = append(errs, err)
		}

		for i := range repos {
			_, err := s.db.SaveRepository(repos[i])
			if err != nil {
				log.Errorf("[user=%s] failed to save [repos=%s] [error=%s]", user.Name, repos[i].Name, err.Error())
				errs = append(errs, fmt.Errorf("failed save [repos=%s] for [user=%s] [error=%w]",
					repos[i].Name, user.Name, err,
				))
			}

			log.Infof("[user=%s] [repos=%s] [remote=%s] saved", user.Name, repos[i].Name, remoteName)
		}

		return combindErr(errs)
	})
	s.userPolls[user.Name] = append(s.userPolls[user.Name], repoPoll)

	return user, err
}

func (s *Service) GetUser(name string) (model.User, error) {
	return s.db.QueryUserByName(name)
}

func (s *Service) GetUserPolls(user model.User) []model.UserPoll {
	var polls []model.UserPoll

	ps := s.userPolls[user.Name]
	for i := range ps {
		polls = append(polls, model.UserPoll{
			ID:     model.UUID(ps[i].ID),
			Name:   user.Name,
			URL:    ps[i].URL,
			Status: ps[i].Status().String(),
		})
	}

	return polls
}
