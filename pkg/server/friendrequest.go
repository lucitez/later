package server

import (
	"later/pkg/transfer"
	"net/http"

	"later/pkg/request"
	"later/pkg/service"

	"github.com/gin-gonic/gin"
)

// FriendRequestServer ...
type FriendRequestServer struct {
	Manager  service.FriendRequest
	Transfer transfer.FriendRequest
}

// NewFriendRequestServer ...
func NewFriendRequestServer(
	manager service.FriendRequest,
	transfer transfer.FriendRequest,
) FriendRequestServer {
	return FriendRequestServer{
		manager,
		transfer,
	}
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *FriendRequestServer) RegisterEndpoints(router *gin.Engine) {
	router.POST("/friend-requests/send", server.send)
	router.GET("/friend-requests/pending", server.pending)
	router.PUT("/friend-requests/accept", server.accept)
	router.PUT("/friend-requests/declilne", server.decline)
}

func (server *FriendRequestServer) send(context *gin.Context) {
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

func (server *FriendRequestServer) pending(context *gin.Context) {
	userID, err := DeserUUID(context, "user_id")

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	friendRequests, err := server.Manager.Pending(*userID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	wireFriendRequests := server.Transfer.WireFriendRequestsFrom(friendRequests)

	context.JSON(http.StatusOK, wireFriendRequests)
}

func (server *FriendRequestServer) accept(context *gin.Context) {
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

func (server *FriendRequestServer) decline(context *gin.Context) {
	var body request.FriendRequestDeclineRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := server.Manager.Decline(body.ID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.Status(http.StatusOK)
}
