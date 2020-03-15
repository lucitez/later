package friendrequestserver

import (
	"net/http"

	"later.co/pkg/request"

	"later.co/pkg/server"

	"later.co/pkg/manager/friendrequestmanager"

	"github.com/gin-gonic/gin"
)

// RegisterEndpoints defines handlers for endpoints for the user service
func RegisterEndpoints(router *gin.Engine) {
	router.POST("/friend-requests/send", send)
	router.GET("/friend-requests/pending", pending)
	router.PUT("/friend-requests/accept", accept)
	router.PUT("/friend-requests/declilne", decline)
}

func send(context *gin.Context) {
	var body request.FriendRequestCreateRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	friendRequest, err := friendrequestmanager.Create(body.ToFriendRequestCreateBody())

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, friendRequest)
}

func pending(context *gin.Context) {
	userID, err := server.DeserUUID(context, "user_id")

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	friendRequest, err := friendrequestmanager.Pending(*userID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, friendRequest)
}

func accept(context *gin.Context) {
	var body request.FriendRequestAcceptRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	err := friendrequestmanager.Accept(body.ID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.Status(http.StatusOK)
}

func decline(context *gin.Context) {
	var body request.FriendRequestDeclineRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	err := friendrequestmanager.Decline(body.ID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.Status(http.StatusOK)
}
