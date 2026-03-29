package service

import (
	"context"
	"errors"
	"fmt"
	"micro/authentication/internal/external"
	"micro/authentication/internal/model"
	storage "micro/authentication/internal/storage/postgresql"
)

type Auth struct {
	stor *storage.Users

	log *external.LogService
}

func NewAuth(stor *storage.Users, log *external.LogService) *Auth {
	return &Auth{
		stor: stor,
		log:  log,
	}
}

func (s *Auth) Authenticate(ctx context.Context, email, password string) (model.User, error) {
	u, err := s.stor.GetUserByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	matched, err := u.PasswordMatches(password)
	if err != nil || !matched {
		return model.User{}, errors.New("invalid credentials")
	}

	_, err = s.log.Log(ctx, model.LogParams{
		Name: "authentication",
		Data: fmt.Sprintf("%s logged in", u.Email),
	})
	if err != nil {
		return model.User{}, err
	}

	return u, nil
}
