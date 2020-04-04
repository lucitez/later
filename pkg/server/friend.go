package server

import (
	"later/pkg/service"
	"later/pkg/transfer"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// FriendServer ...
type FriendServer struct {
	Manager  service.Friend
	Transfer transfer.Friend
}

// NewFriendServer for wire generation
func NewFriendServer(
	manager service.Friend,
	transfer transfer.Friend,
) FriendServer {
	return FriendServer{
		manager,
		transfer,
	}
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *FriendServer) RegisterEndpoints(router *gin.Engine) {
	router.GET("/friends/all", server.all)
	router.GET("/friends/search", server.search)
}

func (server *FriendServer) all(context *gin.Context) {
	deser := NewDeser(
		map[string]Kind{
			"user_id": UUID,
		},
		context,
	)

	if parameters, ok := deser.DeserQueryParams(); ok {
		userID := parameters["user_id"].(uuid.UUID)

		friends := server.Manager.All(userID)
		wireFriends := server.Transfer.WireFriendsFrom(friends)

		context.JSON(http.StatusOK, wireFriends)
	}
}

func (server *FriendServer) search(context *gin.Context) {
	deser := NewDeser(
		map[string]Kind{
			"user_id": UUID,
			"search":  Str,
		},
		context,
	)

	if parameters, ok := deser.DeserQueryParams(); ok {
		userID := parameters["user_id"].(uuid.UUID)
		search := parameters["search"].(string)

		friends := server.Manager.Search(userID, search)
		wireFriends := server.Transfer.WireFriendsFrom(friends)

		context.JSON(http.StatusOK, wireFriends)
	}
}
