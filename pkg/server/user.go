package server

import (
	"net/http"

	"github.com/lucitez/later/pkg/transfer"

	"github.com/lucitez/later/pkg/request"
	"github.com/lucitez/later/pkg/service"

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
		router.GET("/profile", server.profile),                // get own profile
		router.GET("/profile-by-user-id", server.profileByID), // get profile of other user
		router.GET("/search", server.search),

		router.PUT("/update", server.update),
		router.PUT("/update/expo-token", server.updateExpoToken),
	}
}

func (server *User) byID(context *gin.Context) {
	deser := NewDeser(
		context,
		QueryParameter{name: "id", kind: UUID, required: true},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		userID := qp["id"].(*uuid.UUID)
		user, err := server.service.ByID(*userID)

		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if user == nil {
			context.JSON(http.StatusOK, nil)
			return
		}

		context.JSON(http.StatusOK, server.Transfer.WireUserFromUser(*user))
	}
}

func (server *User) profile(context *gin.Context) {
	userID := context.MustGet("user_id").(uuid.UUID)

	user, err := server.service.ByID(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if user == nil {
		context.JSON(http.StatusBadRequest, "User not found")
		return
	}

	wireUser := server.Transfer.WireUserFromUser(*user)

	context.JSON(http.StatusOK, wireUser)
}

func (server *User) profileByID(context *gin.Context) {
	userID := context.MustGet("user_id").(uuid.UUID)

	deser := NewDeser(
		context,
		QueryParameter{name: "profile_user_id", kind: UUID, required: true},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		profileUserID := qp["profile_user_id"].(*uuid.UUID)
		profileUser, err := server.service.ByID(*profileUserID)

		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if profileUser == nil {
			context.JSON(http.StatusBadRequest, "User not found")
			return
		}

		wireUserProfile := server.Transfer.WireUserProfileFrom(userID, *profileUser)

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
		users, err := server.service.Filter(
			search,
			*limit,
			*offset,
		)

		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, server.Transfer.WireUsersFrom(users))
	}
}

func (server *User) updateExpoToken(context *gin.Context) {
	userID := context.MustGet("user_id").(uuid.UUID)

	var body request.UserUpdateExpoToken

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := server.service.UpdateExpoToken(body.Token, userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, true)
}

func (server *User) update(context *gin.Context) {
	userID := context.MustGet("user_id").(uuid.UUID)

	var body request.UserUpdate

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := server.service.Update(body.ToUserUpdateBody(userID))

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, true)
}
