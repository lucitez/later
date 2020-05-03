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
		router.GET("/popular", server.popular),
	}
}

func (server *Content) preview(c *gin.Context) {
	deser := NewDeser(
		c,
		QueryParameter{name: "url", kind: Str, required: true},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		url := qp["url"].(*string)

		contentMetadata, err := server.Service.GetContentPreview(*url)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, contentMetadata.ToContentPreview())
	}
}

func (server *Content) popular(c *gin.Context) {

}
