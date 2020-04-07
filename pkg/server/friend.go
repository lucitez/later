package server

import (
	"later/pkg/service"
	"later/pkg/transfer"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Friend ...
type Friend struct {
	Manager  service.Friend
	Transfer transfer.Friend
}

// NewFriend for wire generation
func NewFriend(
	manager service.Friend,
	transfer transfer.Friend,
) Friend {
	return Friend{
		manager,
		transfer,
	}
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *Friend) RegisterEndpoints(router *gin.Engine) {
	router.GET("/friends/all", server.all)
	router.GET("/friends/search", server.search)
}

func (server *Friend) all(context *gin.Context) {
	deser := NewDeser(
		context,
		QueryParameter{name: "user_id", kind: UUID, required: true},
	)

	if parameters, ok := deser.DeserQueryParams(); ok {
		userID := parameters["user_id"].(uuid.UUID)

		friends := server.Manager.All(userID)
		wireFriends := server.Transfer.WireFriendsFrom(friends)

		context.JSON(http.StatusOK, wireFriends)
	}
}

func (server *Friend) search(context *gin.Context) {
	deser := NewDeser(
		context,
		QueryParameter{name: "user_id", kind: UUID, required: true},
		QueryParameter{name: "search", kind: Str, required: true},
	)

	if parameters, ok := deser.DeserQueryParams(); ok {
		userID := parameters["user_id"].(uuid.UUID)
		search := parameters["search"].(string)

		friends := server.Manager.Search(userID, search)
		wireFriends := server.Transfer.WireFriendsFrom(friends)

		context.JSON(http.StatusOK, wireFriends)
	}
}
