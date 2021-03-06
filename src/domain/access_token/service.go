package access_token

import (
	"strings"

	"github.com/agusluques/bookstore_oauth-api/src/utils/errors"
)

// Repository interface
type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

// Service interface
type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type service struct {
	repository Repository
}

// NewService returns a new access token service
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(id string) (*AccessToken, *errors.RestError) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	return s.repository.GetById(id)
}
