package server

import (
	"net/http"

	"later.co/pkg/manager"

	"github.com/gin-gonic/gin"
	"later.co/pkg/parse"
	"later.co/pkg/request"
)

// ContentServer ...
type ContentServer struct {
	Parser  parse.Parser
	Manager manager.ContentManager
}

// NewContentServer for wire generation
func NewContentServer(
	parser parse.Parser,
	manager manager.ContentManager) ContentServer {
	return ContentServer{
		Parser:  parser,
		Manager: manager}
}

// RegisterEndpoints defines handlers for endpoints for the content service
func (server *ContentServer) RegisterEndpoints(router *gin.Engine) {
	router.POST("/content/create", server.create)
}

func (server *ContentServer) create(context *gin.Context) {
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
