package services

import (
	"strings"

	"github.com/agusluques/bookstore_oauth-api/src/domain/access_token"
	"github.com/agusluques/bookstore_oauth-api/src/repository/db"
	"github.com/agusluques/bookstore_oauth-api/src/repository/rest"
	"github.com/agusluques/bookstore_oauth-api/src/utils/errors"
)

// Repository interface
type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestError)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

// Service interface
type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestError)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DbRepository
}

// NewService returns a new access token service
func NewService(usersRepo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

func (s *service) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(req access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	user, err := s.restUsersRepo.LoginUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	at := access_token.GetNewAccessToken(user.ID)
	at.Generate()

	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
