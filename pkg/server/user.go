package server

import (
	"net/http"
	"strconv"

	"later/pkg/request"
	"later/pkg/service"

	"github.com/gin-gonic/gin"
)

// UserServer ...
type UserServer struct {
	Manager service.User
}

// NewUserServer ...
func NewUserServer(manager service.User) UserServer {
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

	user := server.Manager.SignUp(json)

	context.JSON(http.StatusOK, user)
}

func (server *UserServer) byID(context *gin.Context) {
	userID, err := DeserUUID(context, "user_id")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Parameter id must be a uuid"})
		return
	}

	user := server.Manager.ByID(*userID)

	context.JSON(http.StatusOK, user)
}

func (server *UserServer) allUsers(context *gin.Context) {
	limitstr := context.DefaultQuery("limit", "100")

	limit, err := strconv.Atoi(limitstr)

	users := server.Manager.All(limit)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
}
