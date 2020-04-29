package server

import (
	"later/pkg/request"
	"later/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		router.POST("/create", server.create),
		router.GET("/preview", server.preview),
	}
}

func (server *Content) create(context *gin.Context) {
	userID := context.MustGet("user_id").(uuid.UUID)

	var body request.ContentCreateRequestBody

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	content, err := server.Service.CreateFromURL(body.URL, userID)

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
		url := qp["url"].(*string)

		contentMetadata := server.Service.GetContentPreview(*url)

		context.JSON(http.StatusOK, contentMetadata.ToContentPreview())
	}
}
