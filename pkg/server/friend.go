package server

import (
	"github.com/lucitez/later/pkg/request"
	"github.com/lucitez/later/pkg/service"
	"github.com/lucitez/later/pkg/transfer"
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

func (server *Friend) Prefix() string {
	return "/friends"
}

// Routes defines handlers for endpoints for the user service
func (server *Friend) Routes(router *gin.RouterGroup) []gin.IRoutes {
	return []gin.IRoutes{
		router.GET("/for-user", server.forUser),
		router.PUT("/delete-by-user-id", server.deleteByUserID),
	}
}

func (server *Friend) forUser(context *gin.Context) {
	userID := context.MustGet("user_id").(uuid.UUID)

	defaultLimit := "20"
	defaultOffset := "0"

	deser := NewDeser(
		context,
		QueryParameter{name: "search", kind: Str, required: false},
		QueryParameter{name: "limit", kind: Int, fallback: &defaultLimit},
		QueryParameter{name: "offset", kind: Int, fallback: &defaultOffset},
	)

	if parameters, ok := deser.DeserQueryParams(); ok {
		search := parameters["search"].(*string)
		limit := parameters["limit"].(*int)
		offset := parameters["offset"].(*int)

		friends := server.Service.ForUser(
			userID,
			search,
			*limit,
			*offset,
		)
		wireFriends := server.Transfer.WireFriendsFrom(friends)

		context.JSON(http.StatusOK, wireFriends)
	}
}

func (server *Friend) deleteByUserID(context *gin.Context) {
	userID := context.MustGet("user_id").(uuid.UUID)

	var body request.FriendDeleteRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	server.Service.DeleteByUserID(userID, body.FriendUserID)

	context.JSON(http.StatusOK, true)
}
