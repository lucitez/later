package server

import (
	"github.com/lucitez/later/pkg/service"
	"github.com/lucitez/later/pkg/transfer"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Chat ...
type Chat struct {
	Service  service.Chat
	Transfer transfer.Chat
}

// NewChat for wire generation
func NewChat(
	service service.Chat,
	transfer transfer.Chat,
) Chat {
	return Chat{
		Service:  service,
		Transfer: transfer,
	}
}

func (server *Chat) Prefix() string {
	return "/chats"
}

// Routes defines the routes for chat API
func (server *Chat) Routes(router *gin.RouterGroup) []gin.IRoutes {
	return []gin.IRoutes{
		router.GET("/for-user", server.forUser),
	}
}

func (server *Chat) forUser(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)

	chats, err := server.Service.ForUser(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	wireChats := server.Transfer.WireChatsFrom(chats, userID)

	c.JSON(http.StatusOK, wireChats)
}
