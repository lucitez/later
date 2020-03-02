package usercontentserver

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"later.co/pkg/repository/usercontentrepo"
	"later.co/pkg/util/stringutil"
)

// RegisterEndpoints defines handlers for endpoints for the user service
func RegisterEndpoints(router *gin.Engine) {
	router.GET("/user-content/feed", feed)
}

func feed(context *gin.Context) {

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

	userContent, err := usercontentrepo.Feed(
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
