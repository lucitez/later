package server

import (
	"net/http"
	"strconv"

	"later/pkg/service"
	"later/pkg/request"

	"github.com/gin-gonic/gin"
)

// UserServer ...
type UserServer struct {
	Manager service.UserManager
}

// NewUserServer ...
func NewUserServer(manager service.UserManager) UserServer {
	return UserServer{manager}
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *UserServer) RegisterEndpoints(router *gin.Engine) {
	router.POST("/users/sign-up", server.signUp)

	router.GET("/users/by-id", server.byID)
	router.GET("/users/all", server.allUsers)
}

func (server *UserServer) signUp(context *gin.Context) {
	var json request.UserSignUpRequestBody

	err := context.ShouldBindJSON(&json)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := server.Manager.SignUp(json)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error_type": "On Insert", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (server *UserServer) byID(context *gin.Context) {
	userID, err := DeserUUID(context, "user_id")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Parameter id must be a uuid"})
		return
	}

	user, err := server.Manager.ByID(*userID)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (server *UserServer) allUsers(context *gin.Context) {
	limitstr := context.DefaultQuery("limit", "100")

	limit, err := strconv.Atoi(limitstr)

	users, err := server.Manager.All(limit)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
}
