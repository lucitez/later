package shareserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"later.co/pkg/later/share"
	"later.co/pkg/later/usercontent"
	"later.co/pkg/parse"
	"later.co/pkg/repository/contentrepo"
	"later.co/pkg/repository/sharerepo"
	"later.co/pkg/repository/usercontentrepo"
	"later.co/pkg/request"
)

// RegisterEndpoints defines handlers for endpoints for the user service
func RegisterEndpoints(router *gin.Engine) {
	router.POST("/shares/new", new)
}

/**
*	1. Parse content from URL and create entry in `contents` table
*	2. Insert share into `shares` table
*	3. Create `user_content` from this share.
 */
func new(context *gin.Context) {
	var shareCreateRequestBody request.ShareCreateRequestBody

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
	share, err := share.New(
		content.ID,
		shareCreateRequestBody.SenderUserID,
		shareCreateRequestBody.RecipientUserID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	share, err = sharerepo.Insert(share)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error_type": "On Insert", "error": err.Error()})
		return
	}

	/* create and insert usercontent */
	usercontent, err := usercontent.New(
		share.ID,
		content.ID,
		content.ContentType,
		shareCreateRequestBody.RecipientUserID,
		shareCreateRequestBody.SenderUserID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = usercontentrepo.Insert(usercontent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, share)
}
