package rest

import (
	"encoding/json"
	"time"

	"github.com/agusluques/bookstore_oauth-api/src/domain/users"
	"github.com/agusluques/bookstore_oauth-api/src/utils/errors"
	"github.com/federicoleon/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

type restUsersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &restUsersRepository{}
}

func (r *restUsersRepository) LoginUser(email string, password string) (*users.User, *errors.RestError) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid restClient response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestError
		// should be the same restError interface
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface")
		}

		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshall user response")
	}
	return &user, nil
}
