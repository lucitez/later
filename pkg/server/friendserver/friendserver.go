package friendserver

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// RegisterEndpoints defines handlers for endpoints for the user service
func RegisterEndpoints(router *gin.Engine) {
	router.GET("/friends/all", all)
	router.GET("/friends/search", search)
}

func all(context *gin.Context) {
	userIDStr := context.Query("user_id")

	userID, err := uuid.Parse(userIDStr)

	if err != nil {
		context.JSON(http.StatusBadRequest, "parameter user_id is required and must be UUID")
	}

	context.JSON(http.StatusOK, userID)
}

func search(context *gin.Context) {
	userID := context.MustGet("user_id")
	query := context.MustGet("query")

	fmt.Print(query)

	context.JSON(http.StatusOK, userID)
}
