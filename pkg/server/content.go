package server

import (
	"net/http"

	"github.com/lucitez/later/pkg/service"

	"github.com/gin-gonic/gin"
)

// Content ...
type Content struct {
	Service service.Content
}

// NewContent for wire generation
func NewContent(
	service service.Content,
) Content {
	return Content{
		Service: service,
	}
}

func (server *Content) Prefix() string {
	return "/content"
}

// Routes defines the routes for content API
func (server *Content) Routes(router *gin.RouterGroup) []gin.IRoutes {
	return []gin.IRoutes{
		router.GET("/preview", server.preview),
	}
}

func (server *Content) preview(context *gin.Context) {
	deser := NewDeser(
		context,
		QueryParameter{name: "url", kind: Str, required: true},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		url := qp["url"].(*string)

		contentMetadata, err := server.Service.GetContentPreview(*url)

		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}

		context.JSON(http.StatusOK, contentMetadata.ToContentPreview())
	}
}
