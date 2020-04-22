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

func (server *User) Prefix() string {
	return "/users"
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *User) Routes(router *gin.RouterGroup) []gin.IRoutes {
	return []gin.IRoutes{
		router.GET("/by-id", server.byID),
		router.GET("/profile-by-id", server.profileByID),
		router.GET("/search", server.search),

		router.PUT("/update", server.update),
	}
}

func (server *User) byID(context *gin.Context) {
	deser := NewDeser(
		context,
		QueryParameter{name: "id", kind: UUID, required: true},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		userID := qp["id"].(*uuid.UUID)
		user := server.service.ByID(*userID)

		if user == nil {
			context.JSON(http.StatusOK, nil)
			return
		}

		context.JSON(http.StatusOK, server.Transfer.WireUserFromUser(*user))
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
			return
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

		context.JSON(http.StatusOK, server.Transfer.WireUsersFrom(users))
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
