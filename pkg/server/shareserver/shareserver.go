package shareserver

import (
	"errors"
	"net/http"

	"later.co/pkg/later/entity"
	"later.co/pkg/manager"
	"later.co/pkg/util/wrappers"

	"later.co/pkg/body"

	"github.com/gin-gonic/gin"
	"later.co/pkg/parse"
	"later.co/pkg/request"
)

type ShareServer struct {
	Manager        manager.ShareManager
	ContentManager manager.ContentManager
	UserManager    manager.UserManager
}

// RegisterEndpoints defines handlers for endpoints for the user service
func (server *ShareServer) RegisterEndpoints(router *gin.Engine) {
	router.POST("/shares/new", server.new)
	router.POST("/shares/new/by-phone-number", server.newByPhoneNumber)
}

/**
*	1. If content_id is present, try to get content by that.
*	2. If url is present, parse content from url and insert new content
 */
func (server *ShareServer) getContentFromURLOrContentID(url wrappers.NullString, contentID wrappers.NullUUID) (content *entity.Content, err error) {
	switch {
	case contentID.Valid:
		content, err = server.ContentManager.ByID(contentID.ID)
	case url.Valid:
		contentFromURL, err := parse.ContentFromURL(url.String)
		if err != nil {
			return nil, err
		}
		content, err = server.ContentManager.Create(contentFromURL)
	default:
		content, err = nil, errors.New("parameters url or content_id required")
	}

	return
}

/**
*	1. Get or create _content_ (Get if it is forwarding existing content)
*	2. Create new _share_
*	3. Create new _user_content_ for recipient
 */
func (server *ShareServer) new(context *gin.Context) {
	var requestBody request.ShareCreateRequestBody

	err := context.ShouldBindJSON(&requestBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error_parsing": err.Error()})
		return
	}
	if len(requestBody.RecipientUserIDs) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Must include at least one recipient"})
		return
	}

	content, err := server.getContentFromURLOrContentID(requestBody.URL, requestBody.ContentID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/* create share with inserted content */

	var createBodies []body.ShareCreateBody

	for _, recipientUserID := range requestBody.RecipientUserIDs {
		createBody := body.ShareCreateBody{
			Content:         *content,
			SenderUserID:    requestBody.SenderUserID,
			RecipientUserID: recipientUserID}

		createBodies = append(createBodies, createBody)
	}

	shares, err := server.Manager.CreateMultiple(createBodies)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, shares)
}

/**
*	1. When user wants to share content
 */
func (server *ShareServer) newByPhoneNumber(context *gin.Context) {
	var requestBody request.ShareCreateByPhoneNumberRequestBody

	err := context.ShouldBindJSON(&requestBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userFromPhoneNumber, err := server.userFromPhoneNumber(requestBody.PhoneNumber)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content, err := server.getContentFromURLOrContentID(requestBody.URL, requestBody.ContentID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/* create share with inserted content */

	createBody := body.ShareCreateBody{
		Content:         *content,
		SenderUserID:    requestBody.SenderUserID,
		RecipientUserID: userFromPhoneNumber.ID}

	share, err := server.Manager.Create(createBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, share)
}

/**
*	1. If phone number does belong to an existing user that has signed up, send error response. Client should present an option
*	to add this person as a friend.
*	2. Parse content from URL and create entry in `contents` table
*	3. If phone number does not belong to an existing user or belongs to an existing user that has not signed up,
*	send SMS with URL, Title, and link to us in app store
 */
func (server *ShareServer) userFromPhoneNumber(phoneNumber string) (*entity.User, error) {
	user, err := server.UserManager.ByPhoneNumber(phoneNumber)

	if err != nil {
		return nil, err
	}

	if user != nil && user.SignedUpAt.Valid {
		return nil, errors.New("existing_user_not_friend")
	}

	if user == nil {
		user, err = server.UserManager.NewUserFromPhoneNumber(phoneNumber)
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
