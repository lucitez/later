package server

import (
	"net/http"
	"strconv"

	"later/pkg/service"
	"later/pkg/util/stringutil"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserContentServer ...
type UserContentServer struct {
	Manager service.UserContentManager
}

// NewUserContentServer ...
func NewUserContentServer(manager service.UserContentManager) UserContentServer {
	return UserContentServer{manager}
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *UserContentServer) RegisterEndpoints(router *gin.Engine) {
	router.GET("/user-content/feed", server.feed)
}

func (server *UserContentServer) feed(context *gin.Context) {

	userID := context.Query("user_id")
	senderType := context.Query("sender_type")
	contentType := context.Query("content_type")
	archivedQuery := context.Query("archived")

	archived, err := strconv.ParseBool(archivedQuery)

	if userID == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Parameter user_id is required"})
		return
	}

	userIDAsUUID, err := uuid.Parse(userID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Parameter user_id must be a uuid"})
		return
	}

	userContent, err := server.Manager.Feed(
		userIDAsUUID,
		stringutil.NullIfBlank(&senderType),
		stringutil.NullIfBlank(&contentType),
		&archived)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, userContent)
}
