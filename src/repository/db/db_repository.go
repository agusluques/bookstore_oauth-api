package db

import (
	"github.com/agusluques/bookstore_oauth-api/src/clients/cassandra"
	"github.com/agusluques/bookstore_oauth-api/src/domain/access_token"
	"github.com/agusluques/bookstore_utils-go/rest_errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?"
	queryInsertAccessToken = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?)"
	queryUpdateExpires     = "UPDATE access_tokens SET expires = ? WHERE access_token = ?"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestError)
	Create(access_token.AccessToken) *rest_errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *rest_errors.RestError
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (dbr *dbRepository) GetById(id string) (*access_token.AccessToken, *rest_errors.RestError) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("access token not found")
		}
		return nil, rest_errors.NewInternalServerError(err.Error(), err)
	}

	return &result, nil
}

func (dbr *dbRepository) Create(at access_token.AccessToken) *rest_errors.RestError {
	if err := cassandra.GetSession().Query(queryInsertAccessToken,
		at.AccessToken,
		at.UserID,
		at.ClientID,
		at.Expires,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}
	return nil
}

func (dbr *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *rest_errors.RestError {
	if err := cassandra.GetSession().Query(queryInsertAccessToken,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}
	return nil
}
