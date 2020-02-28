package usercontentserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"later.co/pkg/repository/usercontentrepo"
)

// RegisterEndpoints defines handlers for endpoints for the user service
func RegisterEndpoints(router *gin.Engine) {
	router.GET("/user-content/feed", feed)
}

func feed(context *gin.Context) {

	userID := context.Query("user_id")
	senderType := context.Query("sender_type")
	contentType := context.Query("content_type")

	if userID == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Parameter user_id is required"})
		return
	}

	userIDAsUUID, err := uuid.Parse(userID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Parameter id must be a uuid"})
		return
	}

	userContent, err := usercontentrepo.Feed(
		userIDAsUUID,
		nullIfBlank(&senderType),
		nullIfBlank(&contentType),
		nil)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, userContent)
}

func nullIfBlank(str *string) *string {
	if str != nil && *str == "" {
		return nil
	}
	return str
}
