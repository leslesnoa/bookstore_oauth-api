package access_token_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leslesnoa/bookstore_oauth-api/domain/access_token"
	restErr "github.com/leslesnoa/bookstore_oauth-api/utils/errors"
)

type AccessTokenHandler interface {
	// GetById(string) (*access_token.AccessToken, *restErr.RestErr)
	GetById(*gin.Context)
	Create(*gin.Context)
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

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		resErr := restErr.NewBadRequestError("invalid json body")
		c.JSON(resErr.Status, resErr)
	}

	if err := handler.service.Create(at); err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusCreated, at)
}
