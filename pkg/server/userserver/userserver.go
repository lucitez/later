package userserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"later.co/pkg/later/user"
	"later.co/pkg/repository/userrepo"
	"later.co/pkg/request"
)

// RegisterEndpoints defines handlers for endpoints for the user service
func RegisterEndpoints(router *gin.Engine) {
	router.POST("/users/sign-up", signUp)
	router.GET("/users/by-id", byID)
}

func signUp(context *gin.Context) {
	var json request.UserSignUpRequestBody

	err := context.ShouldBindJSON(&json)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := user.New(
		json.Username,
		json.Email,
		json.PhoneNumber,
		true)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := userrepo.Insert(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error_type": "On Insert", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, createdUser)
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

	user, err := userrepo.ByID(idAsUUID)

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
