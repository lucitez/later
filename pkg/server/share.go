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

func (server *ShareServer) Prefix() string {
	return "shares"
}

// Routes defines handlers for endpoints for the share service
func (server *ShareServer) Routes(router *gin.RouterGroup) []gin.IRoutes {
	return []gin.IRoutes{
		router.POST("/new", server.new),
		router.POST("/forward", server.forward),
		router.POST("/new/by-phone-number", server.newByPhoneNumber),
	}
}

func (server *ShareServer) forward(context *gin.Context) {
	userID := context.MustGet("user_id").(uuid.UUID)

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

	shares, err := server.createSharesFromContent(body.ToShareCreateBodies(userID, *content))

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
	userID := context.MustGet("user_id").(uuid.UUID)

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

	shares, err := server.createSharesFromContent(body.ToShareCreateBodies(userID, *content))

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
	userID := context.MustGet("user_id").(uuid.UUID)

	var requestBody request.ShareCreateByPhoneNumberRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	content, err := server.Content.CreateFromURL(requestBody.URL)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userFromPhoneNumber, err := server.userFromPhoneNumber(requestBody.PhoneNumber)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	createBody := body.ShareCreateBody{
		Content:         *content,
		SenderUserID:    userID,
		RecipientUserID: userFromPhoneNumber.ID,
	}

	share, err := server.Manager.Create(createBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, share)
}

func (server *ShareServer) createSharesFromContent(bodies []body.ShareCreateBody) ([]model.Share, error) {
	var shares []model.Share
	var err error

	for _, body := range bodies {
		share, _ := server.Manager.Create(body)

		if share != nil {
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
