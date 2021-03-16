package rest

import (
	"encoding/json"
	"time"

	"github.com/agusluques/bookstore_oauth-api/src/domain/users"
	"github.com/agusluques/bookstore_utils-go/rest_errors"
	"github.com/federicoleon/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *rest_errors.RestError)
}

type restUsersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &restUsersRepository{}
}

func (r *restUsersRepository) LoginUser(email string, password string) (*users.User, *rest_errors.RestError) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid restClient response when trying to login user", nil)
	}
	if response.StatusCode > 299 {
		var restErr rest_errors.RestError
		// should be the same restError interface
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface", err)
		}

		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshall user response", err)
	}
	return &user, nil
}
