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
	Service service.Content
	Parse   parse.Content
}

// NewContent for wire generation
func NewContent(
	service service.Content,
) Content {
	return Content{
		Service: service,
	}
}

// RegisterEndpoints defines handlers for endpoints for the content service
func (server *Content) RegisterEndpoints(router *gin.Engine) {
	router.POST("/content/create", server.create)
	router.GET("/content/preview", server.preview)
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

func (server *Content) preview(context *gin.Context) {
	deser := NewDeser(
		context,
		QueryParameter{name: "url", kind: Str, required: true},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		domainName := qp["url"].(*string)

		content, err := server.Service.GetContentPreview(*domainName)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, content)
	}
}
