package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// User Binding from json
type User struct {
	ID   uuid.UUID `form:"id" json:"id" binding:"required"`
	Name string    `form:"name" json:"name" binding:"required"`
}

// RegisterEndpoints defines handlers for endpoints for the user service
func RegisterEndpoints(router *gin.Engine) {
	router.POST("/user/create", createUser)
	router.GET("/user/by-id", getUserByID)
}

func createUser(context *gin.Context) {
	var json User

	err := context.ShouldBindJSON(&json)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, json)
}

func getUserByID(context *gin.Context) {
	id := context.DefaultQuery("id", "")

	if id == "" {
		context.String(http.StatusBadRequest, "Parameter id is required\n")
		return
	}

	context.String(http.StatusOK, "Hello %s\n", id)
}
