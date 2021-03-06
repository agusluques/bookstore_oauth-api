package http

import (
	"net/http"

	"github.com/agusluques/bookstore_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
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
