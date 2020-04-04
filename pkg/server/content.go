package server

import (
	"later/pkg/request"
	"later/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Content ...
type Content struct {
	Service service.Content
}

// NewContent for wire generation
func NewContent(service service.Content) Content {
	return Content{Service: service}
}

// RegisterEndpoints defines handlers for endpoints for the content service
func (server *Content) RegisterEndpoints(router *gin.Engine) {
	router.POST("/content/create", server.create)
}

func (server *Content) create(context *gin.Context) {
	var body request.ContentCreateRequestBody

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	content, err := server.Service.CreateFromURL(body.URL)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, content)
}
