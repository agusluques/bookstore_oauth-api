package http

import (
	"net/http"

	"github.com/agusluques/bookstore_oauth-api/src/domain/access_token"
	"github.com/agusluques/bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
	//Update(*gin.Context)
}

type accessTokenHanlder struct {
	service access_token.Service
}

func NewHandler(s access_token.Service) AccessTokenHandler {
	return &accessTokenHanlder{
		service: s,
	}
}

func (ath *accessTokenHanlder) GetById(c *gin.Context) {
	accessToken, err := ath.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (ath *accessTokenHanlder) Create(c *gin.Context) {
	var at access_token.AccessToken
	// TODO: if send an empty body it will not throw error
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := ath.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, at)
}
