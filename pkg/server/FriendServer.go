package server

import (
	"fmt"
	"net/http"

	"later.co/pkg/manager"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type FriendServer struct {
	Manager manager.FriendManager
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *FriendServer) RegisterEndpoints(router *gin.Engine) {
	router.GET("/friends/all", server.all)
	router.GET("/friends/search", server.search)
}

func (server *FriendServer) all(context *gin.Context) {
	userIDStr := context.Query("user_id")

	userID, err := uuid.Parse(userIDStr)

	if err != nil {
		context.JSON(http.StatusBadRequest, "parameter user_id is required and must be UUID")
	}

	context.JSON(http.StatusOK, userID)
}

func (server *FriendServer) search(context *gin.Context) {
	userID := context.MustGet("user_id")
	query := context.MustGet("query")

	fmt.Print(query)

	context.JSON(http.StatusOK, userID)
}
