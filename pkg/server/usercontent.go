package server

import (
	"later/pkg/request"
	"later/pkg/service"
	"later/pkg/transfer"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserContent ...
type UserContent struct {
	Service  service.UserContent
	Transfer transfer.UserContent
}

// NewUserContent ...
func NewUserContent(
	service service.UserContent,
	transfer transfer.UserContent,
) UserContent {
	return UserContent{
		service,
		transfer,
	}
}

func (server *UserContent) Prefix() string {
	return "/user-content"
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *UserContent) Routes(router *gin.RouterGroup) []gin.IRoutes {
	return []gin.IRoutes{
		router.GET("/filter", server.filter),

		router.PUT("/save", server.save),
		router.PUT("/delete", server.delete),
		router.PUT("/update", server.update),
	}
}

func (server *UserContent) filter(context *gin.Context) {
	defaultSaved := "false"
	defaultLimit := "20"

	deser := NewDeser(
		context,
		QueryParameter{name: "user_id", kind: UUID, required: true},
		QueryParameter{name: "tag", kind: Str},
		QueryParameter{name: "content_type", kind: Str},
		QueryParameter{name: "saved", kind: Bool, fallback: &defaultSaved},
		QueryParameter{name: "search", kind: Str},
		QueryParameter{name: "limit", kind: Int, fallback: &defaultLimit},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		userID := qp["user_id"].(*uuid.UUID)
		tag := qp["tag"].(*string)
		contentType := qp["content_type"].(*string)
		saved := qp["saved"].(*bool)
		search := qp["search"].(*string)
		limit := qp["limit"].(*int)

		userContent := server.Service.Filter(
			*userID,
			tag,
			contentType,
			*saved,
			search,
			*limit,
		)

		wireUserContent := server.Transfer.WireUserContentsFrom(userContent)

		context.JSON(http.StatusOK, wireUserContent)
	}
}

func (server *UserContent) save(context *gin.Context) {
	var body request.UserContentSaveRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := server.Service.Save(body.ID, body.Tag); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, true)
}

func (server *UserContent) delete(context *gin.Context) {
	var body request.UserContentDeleteRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := server.Service.Delete(body.ID); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, true)
}

func (server *UserContent) update(context *gin.Context) {
	var body request.UserContentUpdateRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	server.Service.Update(body.ToUserContentUpdateBody())

	context.JSON(http.StatusOK, true)
}
