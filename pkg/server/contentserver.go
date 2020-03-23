package server

import (
	"net/http"

	"later.co/pkg/repository"

	"github.com/gin-gonic/gin"
	"later.co/pkg/parse"
	"later.co/pkg/request"
)

type ContentServer struct {
	Repository repository.ContentRepository
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

	content, err := parse.ContentFromURL(json.URL)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	server.Repository.Insert(content)
}
