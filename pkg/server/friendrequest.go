package server

import (
	"later/pkg/transfer"
	"net/http"

	"later/pkg/request"
	"later/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// FriendRequest ...
type FriendRequest struct {
	Manager  service.FriendRequest
	Transfer transfer.FriendRequest
}

// NewFriendRequest ...
func NewFriendRequest(
	manager service.FriendRequest,
	transfer transfer.FriendRequest,
) FriendRequest {
	return FriendRequest{
		manager,
		transfer,
	}
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *FriendRequest) RegisterEndpoints(router *gin.Engine) {
	router.POST("/friend-requests/send", server.send)
	router.GET("/friend-requests/pending", server.pending)
	router.PUT("/friend-requests/accept", server.accept)
	router.PUT("/friend-requests/declilne", server.decline)
}

func (server *FriendRequest) send(context *gin.Context) {
	var body request.FriendRequestCreateRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	friendRequest, err := server.Manager.Create(body.ToFriendRequestCreateBody())

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, friendRequest)
}

func (server *FriendRequest) pending(context *gin.Context) {
	deser := NewDeser(
		context,
		QueryParameter{"user_id", UUID, nil},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		userID := qp["user_id"].(uuid.UUID)
		friendRequests := server.Manager.Pending(userID)

		wireFriendRequests := server.Transfer.WireFriendRequestsFrom(friendRequests)

		context.JSON(http.StatusOK, wireFriendRequests)
	}
}

func (server *FriendRequest) accept(context *gin.Context) {
	var body request.FriendRequestAcceptRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := server.Manager.Accept(body.ID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

func (server *FriendRequest) decline(context *gin.Context) {
	var body request.FriendRequestDeclineRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	server.Manager.Decline(body.ID)

	context.Status(http.StatusOK)
}
