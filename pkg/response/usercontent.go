package response

import (
	"time"

	"github.com/google/uuid"
	"later/pkg/util/wrappers"
)

// WireUserContent to be sent to client when displaing feed
type WireUserContent struct {
	ID             uuid.UUID
	ContentID      uuid.UUID
	Title          string
	Description    wrappers.NullString
	ImageURL       wrappers.NullString
	ContentType    wrappers.NullString
	Domain         string
	SentBy         uuid.UUID
	SentByUsername wrappers.NullString

	CreatedAt time.Time
}
