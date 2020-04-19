package server

import (
	"later/pkg/transfer"
	"net/http"

	"later/pkg/request"
	"later/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// User ...
type User struct {
	service  service.User
	Transfer transfer.User
}

// NewUser ...
func NewUser(service service.User, transfer transfer.User) User {
	return User{service, transfer}
}

// RegisterEndpoints defines handlers for endpoints for the user service
// TODO add endpoint to get users but exclude users that requester is friends with already
// Same endpoint should return WireFriendUser with information about whether the requesting user has a _pending_ FR
func (server *User) RegisterEndpoints(router *gin.Engine) {
	router.POST("/users/sign-up", server.signUp)

	router.GET("/users/by-id", server.byID)
	router.GET("/users/profile-by-id", server.profileByID)
	router.GET("/users/search", server.search)

	router.PUT("/users/update", server.update)
}

func (server *User) signUp(context *gin.Context) {
	var body request.UserSignUpRequestBody

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := server.service.SignUp(body)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, user)
}

func (server *User) byID(context *gin.Context) {
	deser := NewDeser(
		context,
		QueryParameter{name: "id", kind: UUID, required: true},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		userID := qp["id"].(*uuid.UUID)
		user := server.service.ByID(*userID)

		context.JSON(http.StatusOK, user)
	}
}

func (server *User) profileByID(context *gin.Context) {
	deser := NewDeser(
		context,
		QueryParameter{name: "request_user_id", kind: UUID, required: true},
		QueryParameter{name: "id", kind: UUID, required: true},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		requestUserID := qp["request_user_id"].(*uuid.UUID)
		userID := qp["id"].(*uuid.UUID)
		user := server.service.ByID(*userID)

		if user == nil {
			context.JSON(http.StatusBadRequest, "User not found")
		}

		wireUserProfile := server.Transfer.WireUserProfileFrom(*requestUserID, *user)

		context.JSON(http.StatusOK, wireUserProfile)
	}
}

func (server *User) search(context *gin.Context) {
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
		users := server.service.Filter(
			search,
			*limit,
			*offset,
		)

		context.JSON(http.StatusOK, users)
	}
}

func (server *User) update(context *gin.Context) {
	var body request.UserUpdate

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := server.service.Update(body.ToUserUpdateBody())

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, true)
}
