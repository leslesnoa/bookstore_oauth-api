package access_token_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leslesnoa/bookstore_oauth-api/domain/access_token"
)

type AccessTokenHandler interface {
	// GetById(string) (*access_token.AccessToken, *restErr.RestErr)
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	// accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	// accessToken, err := h.service.GetById(accessTokenId)
	accessToken, err := handler.service.GetById((c.Param("access_token_id")))
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusNotImplemented, accessToken)
}
