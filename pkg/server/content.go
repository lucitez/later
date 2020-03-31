package server

import (
	"later/pkg/parse"
	"later/pkg/request"
	"later/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Content ...
type Content struct {
	Parser  parse.Content
	Manager service.ContentManager
}

// NewContent for wire generation
func NewContent(
	parser parse.Content,
	manager service.ContentManager) Content {
	return Content{
		Parser:  parser,
		Manager: manager}
}

// RegisterEndpoints defines handlers for endpoints for the content service
func (server *Content) RegisterEndpoints(router *gin.Engine) {
	router.POST("/content/create", server.create)
}

func (server *Content) create(context *gin.Context) {
	var json request.ContentCreateRequestBody

	err := context.ShouldBindJSON(&json)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content, err := server.Parser.ContentFromURL(json.URL)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	server.Manager.Create(content)
}
