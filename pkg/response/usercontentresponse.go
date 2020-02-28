package response

import (
	"time"

	"github.com/google/uuid"
	"later.co/pkg/util/wrappers"
)

// WireUserContent to be sent to client when displaing feed
type WireUserContent struct {
	ID             uuid.UUID
	Title          string
	Description    wrappers.NullString
	ImageURL       wrappers.NullString
	ContentType    wrappers.NullString
	Domain         string
	SentBy         uuid.UUID
	SentByUsername string

	SentAt time.Time
}
