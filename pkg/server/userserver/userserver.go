package userserver

import (
	"net/http"
	"strconv"

	"later.co/pkg/later/entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"later.co/pkg/repository/userrepo"
	"later.co/pkg/request"
)

// RegisterEndpoints defines handlers for endpoints for the user service
func RegisterEndpoints(router *gin.Engine) {
	router.POST("/users/sign-up", signUp)

	router.GET("/users/by-id", byID)
	router.GET("/users/all", allUsers)
}

func signUp(context *gin.Context) {
	var json request.UserSignUpRequestBody

	err := context.ShouldBindJSON(&json)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := entity.NewUser(
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
		context.JSON(http.StatusBadRequest, gin.H{"error": "Parameter id is required"})
		return
	}

	idAsUUID, err := uuid.Parse(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Parameter id must be a uuid"})
		return
	}

	user, err := userrepo.ByID(idAsUUID)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func allUsers(context *gin.Context) {
	limitstr := context.DefaultQuery("limit", "100")

	limit, err := strconv.Atoi(limitstr)

	users, err := userrepo.All(limit)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
}
