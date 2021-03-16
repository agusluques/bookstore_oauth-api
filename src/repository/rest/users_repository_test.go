package rest

import (
	"net/http"
	"os"
	"testing"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8081/users/login",
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: -1,
	})

	repository := restUsersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restClient response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8081/users/login",
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": 1}`,
	})

	repository := restUsersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8081/users/login",
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid user credential", "status": 404, "error": "not_found"}`,
	})

	repository := restUsersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid user credential", err.Message)
}

func TestLoginUserInvalidJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8081/users/login",
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "invalid id", "first_name":"Test", "last_name":"Test", "email":"email@gmail.com"}`,
	})

	repository := restUsersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshall user response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8081/users/login",
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":1, "first_name":"Test", "last_name":"Test", "email":"email@gmail.com"}`,
	})

	repository := restUsersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.NotNil(t, user)
	assert.Nil(t, err)
}
