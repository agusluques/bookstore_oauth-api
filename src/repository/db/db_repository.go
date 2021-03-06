package db

import (
	"github.com/agusluques/bookstore_oauth-api/src/domain/access_token"
	"github.com/agusluques/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (dbr *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	return nil, errors.NewInternalServerError("db not implemented")
}
