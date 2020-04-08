package server

import (
	"net/http"

	"later/pkg/request"
	"later/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// User ...
type User struct {
	Manager service.User
}

// NewUser ...
func NewUser(manager service.User) User {
	return User{manager}
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *User) RegisterEndpoints(router *gin.Engine) {
	router.POST("/users/sign-up", server.signUp)

	router.GET("/users/by-id", server.byID)
	router.GET("/users/all", server.allUsers)
}

func (server *User) signUp(context *gin.Context) {
	var body request.UserSignUpRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := server.Manager.SignUp(body)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, user)
}

func (server *User) byID(context *gin.Context) {
	deser := NewDeser(
		context,
		QueryParameter{name: "id", kind: UUID},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		userID := qp["id"].(*uuid.UUID)
		user := server.Manager.ByID(*userID)

		context.JSON(http.StatusOK, user)
	}
}

func (server *User) allUsers(context *gin.Context) {
	defaultLimit := "100"

	deser := NewDeser(
		context,
		QueryParameter{name: "limit", kind: Int, fallback: &defaultLimit},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		limit := qp["limit"].(*int)
		users := server.Manager.All(*limit)

		context.JSON(http.StatusOK, users)
	}
}
