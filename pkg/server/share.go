package server

import (
	"errors"
	"net/http"

	"github.com/google/uuid"

	"later/pkg/model"
	"later/pkg/parse"
	"later/pkg/service"
	"later/pkg/service/body"

	"later/pkg/request"

	"github.com/gin-gonic/gin"
)

// ShareServer ...
type ShareServer struct {
	Manager service.Share
	Content service.Content
	User    service.User
	Parser  parse.Content
}

// NewShareServer ...
func NewShareServer(
	manager service.Share,
	contentManager service.Content,
	userManager service.User,
	parser parse.Content,
) ShareServer {
	return ShareServer{
		Manager: manager,
		Content: contentManager,
		User:    userManager,
		Parser:  parser,
	}
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *ShareServer) RegisterEndpoints(router *gin.Engine) {
	router.POST("/shares/new", server.new)
	router.POST("/shares/forward", server.new)
	router.POST("/shares/new/by-phone-number", server.newByPhoneNumber)
}

func (server *ShareServer) forward(context *gin.Context) {
	var body request.ShareForwardRequestBody

	err := context.ShouldBindJSON(&body)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	content := server.Content.ByID(body.ContentID)

	if content == nil {
		context.JSON(http.StatusBadRequest, errors.New("Content not found"))
		return
	}

	shares, err := server.createSharesFromContent(*content, body.SenderUserID, body.RecipientUserIDs)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, shares)
}

/**
*	1. Get or create _content_ (Get if it is forwarding existing content)
*	2. Create new _share_
*	3. Create new _user_content_ for recipient
 */
func (server *ShareServer) new(context *gin.Context) {
	var body request.ShareCreateRequestBody

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	content, err := server.Content.CreateFromURL(body.URL)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	shares, err := server.createSharesFromContent(*content, body.SenderUserID, body.RecipientUserIDs)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, shares)
}

/**
*	1. When user wants to share content
 */
func (server *ShareServer) newByPhoneNumber(context *gin.Context) {
	var body request.ShareCreateByPhoneNumberRequestBody

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	content, err := server.Content.CreateFromURL(body.URL)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userFromPhoneNumber, err := server.userFromPhoneNumber(body.PhoneNumber)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	shares, err := server.createSharesFromContent(*content, body.SenderUserID, []uuid.UUID{userFromPhoneNumber.ID})

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, shares)
}

func (server *ShareServer) createSharesFromContent(content model.Content, sharer uuid.UUID, sharees []uuid.UUID) ([]model.Share, error) {
	var shares []model.Share
	var err error

	for _, sharee := range sharees {
		createBody := body.ShareCreateBody{
			Content:         content,
			SenderUserID:    sharer,
			RecipientUserID: sharee,
		}

		share, err := server.Manager.Create(createBody)

		if err != nil {
			break
		} else if share != nil {
			shares = append(shares, *share)
		}
	}

	return shares, err
}

/**
*	1. If phone number does belong to an existing user that has signed up, send error response. Client should present an option
*	to add this person as a friend.
*	2. Parse content from URL and create entry in `contents` table
*	3. If phone number does not belong to an existing user or belongs to an existing user that has not signed up,
*	send SMS with URL, Title, and link to us in app store
 */
func (server *ShareServer) userFromPhoneNumber(phoneNumber string) (*model.User, error) {
	user := server.User.ByPhoneNumber(phoneNumber)

	if user != nil {
		if user.SignedUpAt.Valid {
			return nil, errors.New("existing_user_not_friend")
		}
		return user, nil
	}

	return server.User.NewUserFromPhoneNumber(phoneNumber)
}
