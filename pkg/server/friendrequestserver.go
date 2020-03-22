package server

import (
	"net/http"

	"later.co/pkg/request"

	"later.co/pkg/manager"

	"github.com/gin-gonic/gin"
)

type FriendRequestServer struct {
	Manager manager.FriendRequestManager
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *FriendRequestServer) RegisterEndpoints(router *gin.Engine) {
	router.POST("/friend-requests/send", server.send)
	router.GET("/friend-requests/pending", server.pending)
	router.PUT("/friend-requests/accept", server.accept)
	router.PUT("/friend-requests/declilne", server.decline)
}

func (frServer *FriendRequestServer) send(context *gin.Context) {
	var body request.FriendRequestCreateRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	friendRequest, err := frServer.Manager.Create(body.ToFriendRequestCreateBody())

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, friendRequest)
}

func (frServer *FriendRequestServer) pending(context *gin.Context) {
	userID, err := DeserUUID(context, "user_id")

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	friendRequest, err := frServer.Manager.Pending(*userID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, friendRequest)
}

func (frServer *FriendRequestServer) accept(context *gin.Context) {
	var body request.FriendRequestAcceptRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	err := frServer.Manager.Accept(body.ID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.Status(http.StatusOK)
}

func (frServer *FriendRequestServer) decline(context *gin.Context) {
	var body request.FriendRequestDeclineRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	err := frServer.Manager.Decline(body.ID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.Status(http.StatusOK)
}
