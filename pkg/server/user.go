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
// TODO add endpoint to get users but exclude users that requester is friends with already
// Same endpoint should return WireFriendUser with information about whether the requesting user has a _pending_ FR
func (server *User) RegisterEndpoints(router *gin.Engine) {
	router.POST("/users/sign-up", server.signUp)

	router.GET("/users/by-id", server.byID)
	router.GET("/users/filter", server.filter)
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

func (server *User) filter(context *gin.Context) {
	defaultLimit := "20"
	defaultOffset := "0"

	deser := NewDeser(
		context,
		QueryParameter{name: "search", kind: Str, required: false},
		QueryParameter{name: "limit", kind: Int, fallback: &defaultLimit},
		QueryParameter{name: "offset", kind: Int, fallback: &defaultOffset},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		search := qp["search"].(*string)
		limit := qp["limit"].(*int)
		offset := qp["offset"].(*int)
		users := server.Manager.Filter(
			search,
			*limit,
			*offset,
		)

		context.JSON(http.StatusOK, users)
	}
}
