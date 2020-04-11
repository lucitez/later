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
	router.GET("/friends/for-user", server.forUser)
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

		friends := server.Manager.ForUser(
			*userID,
			search,
			*limit,
			*offset,
		)
		wireFriends := server.Transfer.WireFriendsFrom(friends)

		context.JSON(http.StatusOK, wireFriends)
	}
}
