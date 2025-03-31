package service

import (
	"ticketsbooking/libs/entities"
	"ticketsbooking/services/auth/internal/interfaces"
	"ticketsbooking/services/auth/internal/privateauth"

	"golang.org/x/crypto/bcrypt"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type service struct {
	repo  interfaces.Repository
	cache interfaces.CacheDB
}

func New(repo interfaces.Repository, cache interfaces.CacheDB) interfaces.Service {
	return &service{repo: repo, cache: cache}
}

func (svc *service) Register(user entities.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "generating hashed password")
	}
	user.Password = string(hashedPass)

	err = svc.repo.CreateUser(user)
	if err != nil {
		return err
	}

	err = svc.cache.SetUser(user)
	return err
}

func (svc *service) Login(user entities.User) (string, error) {
	var foundUser entities.User
	foundUser, err := svc.cache.GetUser(user)
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			foundUser, err = svc.repo.GetUserByEmail(user)
			if err != nil {
				return "", err
			}

			err = svc.cache.SetUser(foundUser)
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		return "", errors.Wrap(err, "comparing hash and password")
	}

	token, err := privateauth.CreateToken(foundUser)
	if err != nil {
		return "", err
	}

	return token, nil
}
