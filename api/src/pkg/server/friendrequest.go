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

func (server *FriendRequest) Prefix() string {
	return "/friend-requests"
}

// Routes defines handlers for endpoints for the friend request service
func (server *FriendRequest) Routes(router *gin.RouterGroup) []gin.IRoutes {
	return []gin.IRoutes{
		router.POST("/send", server.send),

		router.GET("/pending", server.pending),

		router.PUT("/delete", server.delete),
		router.PUT("/accept", server.accept),
		router.PUT("/decline", server.decline),
	}
}

func (server *FriendRequest) send(context *gin.Context) {
	userID := context.MustGet("user_id").(uuid.UUID)

	var body request.FriendRequestCreateRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	friendRequest, err := server.Manager.Create(body.ToFriendRequestCreateBody(userID))

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, friendRequest)
}

func (server *FriendRequest) pending(context *gin.Context) {
	userID := context.MustGet("user_id").(uuid.UUID)

	friendRequests := server.Manager.Pending(userID)
	wireFriendRequests := server.Transfer.WireFriendRequestsFrom(friendRequests)

	context.JSON(http.StatusOK, wireFriendRequests)
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

func (server *FriendRequest) delete(context *gin.Context) {
	var body request.FriendRequestDeleteRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	server.Manager.Delete(body.ID)

	context.Status(http.StatusOK)
}
