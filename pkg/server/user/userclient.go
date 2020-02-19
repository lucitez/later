package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"later.co/pkg/later"
	userRepository "later.co/pkg/repository/user"
)

// RegisterEndpoints defines handlers for endpoints for the user service
func RegisterEndpoints(router *gin.Engine) {
	router.POST("/user/sign-up", signUp)
	router.GET("/user/by-id", byID)
}

func signUp(context *gin.Context) {
	var json later.UserSignUpRequestBody

	err := context.ShouldBindJSON(&json)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, json)
}

func byID(context *gin.Context) {
	id := context.Query("id")

	if id == "" {
		context.String(http.StatusBadRequest, "Parameter id is required\n")
		return
	}

	idAsUUID, err := uuid.Parse(id)

	if err != nil {
		context.String(http.StatusBadRequest, "Parameter id must be a uuid\n")
		return
	}

	user, err := userRepository.ByID(idAsUUID)

	if err != nil {
		fmt.Println(err)
		context.String(http.StatusNotFound, "User does not exist\n")
		return
	}

	response, err := json.Marshal(user)

	if err != nil {
		context.String(http.StatusInternalServerError, "Error converting object into json")
	}

	context.String(http.StatusOK, string(response))
}
