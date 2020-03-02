package shareserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"later.co/pkg/later/share"
	"later.co/pkg/manager/sharemanager"
	"later.co/pkg/parse"
	"later.co/pkg/repository/contentrepo"
	"later.co/pkg/request"
)

// RegisterEndpoints defines handlers for endpoints for the user service
func RegisterEndpoints(router *gin.Engine) {
	router.POST("/shares/forward", forwardShare)
	router.POST("/shares/new/by-user-id", newByUserID)
	router.POST("/shares/new/by-phone-number", newByPhoneNumber)
}

/**
*	1. Already have content from the share reference
*	For each user in recipients list:
*	2. Create new share
*	3. Create new user_content
*	4. Send notification
 */
func forwardShare(context *gin.Context) {

}

/**
*	1. If phone number does belong to an existing user, send error response. Client should present an option
*	to add this person as a friend.
*	2. Parse content from URL and create entry in `contents` table
*	3. If phone number does not belong to an existing user, send SMS with URL, Title, and link to us in app store
*	4.
 */
func newByPhoneNumber(context *gin.Context) {

}

/**
*	1. Parse content from URL and create entry in `contents` table
*	2. Insert share into `shares` table
*	3. Create `user_content` from this share.
 */
func newByUserID(context *gin.Context) {
	var shareCreateRequestBody request.ShareCreateByUserIDRequestBody

	err := context.ShouldBindJSON(&shareCreateRequestBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/* get and insert content from url */
	content, err := parse.ContentFromURL(shareCreateRequestBody.URL)

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

	createBodies := shareCreateRequestBody.ToShareCreateBodies(content)

	shares := []share.Share{}

	for _, createBody := range createBodies {
		share, err := sharemanager.Create(createBody)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		shares = append(shares, *share)
	}

	context.JSON(http.StatusOK, shares)
}
