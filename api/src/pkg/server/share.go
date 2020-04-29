package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/lucitez/later/api/src/pkg/model"
	"github.com/lucitez/later/api/src/pkg/service"
	"github.com/lucitez/later/api/src/pkg/service/body"

	"github.com/lucitez/later/api/src/pkg/request"

	"github.com/gin-gonic/gin"
)

// ShareServer ...
type ShareServer struct {
	Manager service.Share
	Content service.Content
	User    service.User
}

// NewShareServer ...
func NewShareServer(
	manager service.Share,
	contentManager service.Content,
	userManager service.User,
) ShareServer {
	return ShareServer{
		Manager: manager,
		Content: contentManager,
		User:    userManager,
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

	content, _ := server.Content.ByID(body.ContentID)

	if content == nil {
		context.JSON(http.StatusBadRequest, errors.New("Content not found"))
		return
	}

	shares := server.createSharesFromContent(body.ToShareCreateBodies(userID, *content))

	go server.Content.IncrementShareCount(content.ID, len(shares))

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

	content, err := server.Content.CreateFromURL(body.URL, userID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	shares := server.createSharesFromContent(body.ToShareCreateBodies(userID, *content))

	context.JSON(http.StatusOK, shares)
}

func (server *ShareServer) createSharesFromContent(bodies []body.ShareCreateBody) []model.Share {
	var shares []model.Share

	for _, body := range bodies {
		share, err := server.Manager.Create(body)

		if err != nil {
			log.Printf("[WARN] Error creating share %v", body)
		}

		if share != nil {
			shares = append(shares, *share)
		}
	}

	return shares
}
