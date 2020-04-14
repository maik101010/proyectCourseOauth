package http

import (
	"github.com/gin-gonic/gin"
	accessTokenDomain "github.com/maik101010/proyectCourseOauth/src/domain/access_token"
	"github.com/maik101010/proyectCourseOauth/src/service/access_token"
	"github.com/maik101010/proyectCourseOauth/src/util/errors"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}
func NewHandler(service access_token.Service) AccessTokenHandler  {
	return &accessTokenHandler{
		service: service,
	}
}
func (h *accessTokenHandler) GetById(c *gin.Context){
	accessToken, err := h.service.GetById(c.Param("access_token_id"))
	if err!=nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
func (h *accessTokenHandler) Create(c *gin.Context){
	var accessTokenRequest accessTokenDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&accessTokenRequest); err!=nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}
	accessToken, err := h.service.Create(accessTokenRequest)
	if err!=nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}

