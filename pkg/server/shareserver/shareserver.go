package shareserver

import (
	"errors"
	"net/http"

	"later.co/pkg/body"

	"github.com/gin-gonic/gin"
	"later.co/pkg/later/user"
	"later.co/pkg/manager/sharemanager"
	"later.co/pkg/manager/usermanager"
	"later.co/pkg/parse"
	"later.co/pkg/repository/contentrepo"
	"later.co/pkg/repository/userrepo"
	"later.co/pkg/request"
)

// RegisterEndpoints defines handlers for endpoints for the user service
func RegisterEndpoints(router *gin.Engine) {
	router.POST("/shares/forward", forward)
	router.POST("/shares/new", new)
}

/**
*	1. Already have content from the share reference
*	For each user in recipients list:
*	2. Create new share
*	3. Create new user_content
*	4. Send notification
 */
func forward(context *gin.Context) {
	var requestBody request.ShareForwardRequestBody

	err := context.ShouldBindJSON(&requestBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIds := requestBody.RecipientUserIDs

	if requestBody.PhoneNumber.Valid {
		userFromPhoneNumber, err := userFromPhoneNumber(requestBody.PhoneNumber.String)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userIds = append(userIds, userFromPhoneNumber.ID)
	}

	content, err := contentrepo.ByID(requestBody.ContentID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if content == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Content not found"})
		return
	}

	var createBodies []body.ShareCreateBody

	for _, recipientUserID := range userIds {
		createBody := body.ShareCreateBody{
			Content:         *content,
			SenderUserID:    requestBody.SenderUserID,
			RecipientUserID: recipientUserID}

		createBodies = append(createBodies, createBody)
	}

	shares, err := sharemanager.CreateMultiple(createBodies)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, shares)
}

/**
*	1. Parse content from URL and create entry in `contents` table
*	2. Insert share into `shares` table
*	3. Create `user_content` from this share.
 */
func new(context *gin.Context) {
	var requestBody request.ShareCreateRequestBody

	err := context.ShouldBindJSON(&requestBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIds := requestBody.RecipientUserIDs

	if requestBody.PhoneNumber.Valid {
		userFromPhoneNumber, err := userFromPhoneNumber(requestBody.PhoneNumber.String)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userIds = append(userIds, userFromPhoneNumber.ID)
	}

	/* get and insert content from url */
	content, err := parse.ContentFromURL(requestBody.URL)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content, err = contentrepo.Insert(content)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/* create share with inserted content */

	var createBodies []body.ShareCreateBody

	for _, recipientUserID := range userIds {
		createBody := body.ShareCreateBody{
			Content:         *content,
			SenderUserID:    requestBody.SenderUserID,
			RecipientUserID: recipientUserID}

		createBodies = append(createBodies, createBody)
	}

	shares, err := sharemanager.CreateMultiple(createBodies)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, shares)
}

/**
*	1. If phone number does belong to an existing user that has signed up, send error response. Client should present an option
*	to add this person as a friend.
*	2. Parse content from URL and create entry in `contents` table
*	3. If phone number does not belong to an existing user or belongs to an existing user that has not signed up,
*	send SMS with URL, Title, and link to us in app store
 */
func userFromPhoneNumber(phoneNumber string) (*user.User, error) {
	user, err := userrepo.ByPhoneNumber(phoneNumber)

	if err != nil {
		return nil, err
	}

	if user != nil && user.SignedUpAt.Valid {
		return nil, errors.New("existing_user_not_friend")
	}

	if user == nil {
		user, err = usermanager.NewUserFromPhoneNumber(phoneNumber)
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
