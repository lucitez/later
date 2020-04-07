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

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *UserContent) RegisterEndpoints(router *gin.Engine) {
	router.GET("/user-content/filter", server.filter)

	router.PUT("/user-content/archive", server.archive)
	router.PUT("/user-content/delete", server.delete)
}

func (server *UserContent) filter(context *gin.Context) {
	defaultArchived := "false"
	defaultLimit := "20"

	deser := NewDeser(
		context,
		QueryParameter{name: "user_id", kind: UUID, required: true},
		QueryParameter{name: "tag", kind: Str},
		QueryParameter{name: "content_type", kind: Str},
		QueryParameter{name: "archived", kind: Bool, fallback: &defaultArchived},
		QueryParameter{name: "limit", kind: Int, fallback: &defaultLimit},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		userID := qp["user_id"].(*uuid.UUID)
		tag := qp["tag"].(*string)
		contentType := qp["content_type"].(*string)
		archived := qp["archived"].(*bool)
		limit := qp["limit"].(*int)

		userContent := server.Service.Filter(
			*userID,
			tag,
			contentType,
			*archived,
			*limit,
		)

		wireUserContent := server.Transfer.WireUserContentsFrom(userContent)

		context.JSON(http.StatusOK, wireUserContent)
	}
}

func (server *UserContent) archive(context *gin.Context) {
	var body request.UserContentArchiveRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := server.Service.Archive(body.ID, body.Tag); err != nil {
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
