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
	LocalCacheMap map[string]string

	db        db.DB
	userPolls map[string][]*Poll
}

func NewService(db db.DB) *Service {
	remote.Setup(db)

	return &Service{
		LocalCacheMap: make(map[string]string),

		db:        db,
		userPolls: make(map[string][]*Poll),
	}
}

func (s *Service) RegisterUser(remoteName string, details model.UserGitRegisterDetails) (model.User, error) {
	user := model.User{Tokens: make(map[string]string)}

	remoteService := remote.GetRemoteService(remoteName)
	token, err := remoteService.FetchNewToken(details)
	if err != nil {
		log.Errorf("failed to fetch new token from [state=%s] [error=%s]", details.State, err.Error())
		return user, err
	}

	user, err = remoteService.RegisterUser(token)
	if err != nil {
		log.Errorf("[user=%s] failed to register with [remote=%s]", user.Name, remoteName)
		return user, err
	}
	log.Infof("[user=%s] registered with [remote=%s]", user.Name, remoteName)

	tmpUser, err := s.db.QueryUserByName(user.Name)
	if err != nil {
		log.Errorf("[user=%s] failed to query user [error=%s]", user.Name, err.Error())
		return user, err
	}

	if tmpUser.Name == user.Name {
		log.Debugf("[user=%s] already exists", user.Name)
		tmpUser.Tokens[remoteName] = token
		return tmpUser, nil
	}

	// TODO: think about this more...
	// TODO: should we allow other users to exists
	// TODO: and only allow them to view owner's orgs if they are a member?
	if userCount, err := s.db.QueryUserCount(); userCount != 0 || err != nil {
		log.Debugf("only allowing owner to register [user=%s]", user.Name)
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

		orgs, err := remoteService.RegisterUserOrganizations(user.Tokens[remoteName], user)
		if err != nil {
			log.Errorf("[user=%s] failed to register orgs [error=%s]", user.Name, err.Error())
			errs = append(errs, fmt.Errorf("failed to register orgs for [user=%s] [error=%w]", user.Name, err))
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

		log.Infof("[user=%s] saved/updated all orgs [remote=%s]", user.Name, remoteName)
		return combindErr(errs)
	})
	s.userPolls[user.Name] = append(s.userPolls[user.Name], orgPoll)

	repoPoll := NewPoll(context.Background(), "/repos", func(ctx context.Context, p *Poll) error {
		var errs []error

		repos, err := remoteService.RegisterUserRepositories(user.Tokens[remoteName], user)
		if err != nil {
			log.Errorf("[user=%s] failed to register repos [error=%s]", user.Name, err.Error())
			errs = append(errs, fmt.Errorf("failed to register repos for [user=%s] [error=%w]", user.Name, err))
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

		log.Infof("[user=%s] saved/updated all repos [remote=%s]", user.Name, remoteName)
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
			UUID:   model.UUID(ps[i].ID),
			Name:   user.Name,
			URL:    ps[i].URL,
			Status: ps[i].Status().String(),
		})
	}

	return polls
}
