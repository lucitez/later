package server

import (
	"net/http"

	"later.co/pkg/manager"

	"github.com/gin-gonic/gin"
	"later.co/pkg/parse"
	"later.co/pkg/request"
)

type ContentServer struct {
	Router  *gin.Engine
	Parser  parse.Parser
	Manager manager.ContentManager
}

func NewContentServer(
	parser parse.Parser,
	manager manager.ContentManager) ContentServer {
	return ContentServer{
		Parser:  parser,
		Manager: manager}
}

func (server *ContentServer) Start() {
	server.RegisterEndpoints()
}

// RegisterEndpoints defines handlers for endpoints for the content service
func (server *ContentServer) RegisterEndpoints() {
	server.Router.POST("/content/create", server.create)
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
