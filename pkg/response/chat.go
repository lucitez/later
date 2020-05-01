package response

import (
	"github.com/lucitez/later/pkg/util/wrappers"

	"github.com/google/uuid"
)

type WireChat struct {
	ChatID            uuid.UUID
	OtherUserID       wrappers.NullUUID
	OtherUserUsername wrappers.NullString
	OtherUserName     wrappers.NullString
	GroupName         wrappers.NullString
	Activity          string
	LastMessageSentAt wrappers.NullTime
	HasUnread         bool
}
