package transfer

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/response"
	"github.com/lucitez/later/pkg/service"
)

type UserContent struct {
	ContentService service.Content
	UserService    service.User
}

func NewUserContent(
	contentService service.Content,
	userService service.User,
) UserContent {
	return UserContent{
		contentService,
		userService,
	}
}

// WireUserContentsFrom tranfers DB model UserContent to DTO WireUserContent
func (transfer *UserContent) WireUserContentsFrom(userContent []model.UserContent) (wire []response.WireUserContent, err error) {
	wire = make([]response.WireUserContent, len(userContent))

	var content *model.Content

	for i, userContent := range userContent {
		content, err = transfer.ContentService.ByID(userContent.ContentID)

		if err != nil {
			return
		}

		user, _ := transfer.UserService.ByID(userContent.SentByUserID)
		if content != nil && user != nil {
			wire[i] = wireUserContent(userContent, *content, *user)
		}
	}

	return
}

func wireUserContent(userContent model.UserContent, content model.Content, user model.User) response.WireUserContent {
	return response.WireUserContent{
		ID:             userContent.ID,
		ContentID:      content.ID,
		Title:          content.Title,
		Description:    content.Description,
		ImageURL:       content.ImageURL,
		ContentType:    content.ContentType,
		Tag:            userContent.Tag,
		Hostname:       content.Hostname,
		URL:            content.URL,
		Shares:         content.Shares,
		SentBy:         userContent.SentByUserID,
		SentByUsername: user.Username,
		SentByName:     user.Name,
		CreatedAt:      userContent.CreatedAt,
		SavedAt:        userContent.SavedAt,
	}
}
