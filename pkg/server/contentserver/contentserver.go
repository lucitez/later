package contentserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"later.co/pkg/parse"
	"later.co/pkg/request"
)

// RegisterEndpoints defines handlers for endpoints for the content service
func RegisterEndpoints(router *gin.Engine) {
	router.POST("/content/create", create)
}

func create(context *gin.Context) {
	var json request.ContentCreateRequestBody

	err := context.ShouldBindJSON(&json)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = parse.ContentFromURL(json.URL)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
