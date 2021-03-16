package http

import (
	"net/http"

	"github.com/agusluques/bookstore_oauth-api/src/domain/access_token"
	"github.com/agusluques/bookstore_oauth-api/src/services"
	"github.com/agusluques/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service services.Service
}

func NewHandler(s services.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: s,
	}
}

func (ath *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := ath.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (ath *accessTokenHandler) Create(c *gin.Context) {
	var request access_token.AccessTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := ath.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
