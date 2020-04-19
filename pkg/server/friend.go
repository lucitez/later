package server

import (
	"later/pkg/request"
	"later/pkg/service"
	"later/pkg/transfer"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Friend ...
type Friend struct {
	Service  service.Friend
	Transfer transfer.Friend
}

// NewFriend for wire generation
func NewFriend(
	service service.Friend,
	transfer transfer.Friend,
) Friend {
	return Friend{
		service,
		transfer,
	}
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *Friend) RegisterEndpoints(router *gin.Engine) {
	router.GET("/friends/for-user", server.forUser)
	router.PUT("/friends/delete-by-user-id", server.deleteByUserID)
}

func (server *Friend) forUser(context *gin.Context) {
	defaultLimit := "20"
	defaultOffset := "0"

	deser := NewDeser(
		context,
		QueryParameter{name: "user_id", kind: UUID, required: true},
		QueryParameter{name: "search", kind: Str, required: false},
		QueryParameter{name: "limit", kind: Int, fallback: &defaultLimit},
		QueryParameter{name: "offset", kind: Int, fallback: &defaultOffset},
	)

	if parameters, ok := deser.DeserQueryParams(); ok {
		userID := parameters["user_id"].(*uuid.UUID)
		search := parameters["search"].(*string)
		limit := parameters["limit"].(*int)
		offset := parameters["offset"].(*int)

		friends := server.Service.ForUser(
			*userID,
			search,
			*limit,
			*offset,
		)
		wireFriends := server.Transfer.WireFriendsFrom(friends)

		context.JSON(http.StatusOK, wireFriends)
	}
}

func (server *Friend) deleteByUserID(context *gin.Context) {
	var body request.FriendDeleteRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	server.Service.DeleteByUserID(body.UserID, body.FriendUserID)

	context.JSON(http.StatusOK, true)
}
