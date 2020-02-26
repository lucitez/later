package shareserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"later.co/pkg/later/share"
	"later.co/pkg/parse"
	"later.co/pkg/repository/contentrepo"
	"later.co/pkg/repository/sharerepo"
	"later.co/pkg/request"
)

// RegisterEndpoints defines handlers for endpoints for the user service
func RegisterEndpoints(router *gin.Engine) {
	router.POST("/shares/new", new)
}

/**
*	1. Parse content from URL and create entry in `contents` table
*	2. Insert share into `shares` table
*	3. TODO Create `user_content` from this share.
 */
func new(context *gin.Context) {
	var json request.ShareCreateRequestBody

	err := context.ShouldBindJSON(&json)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/* get and insert content from url */
	content, err := parse.ContentFromURL(json.URL)

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
		json.SenderUserID,
		json.RecipientUserID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	share, err = sharerepo.Insert(share)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error_type": "On Insert", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, share)
}
