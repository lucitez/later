package response

import (
	"github.com/google/uuid"
)

type WireChat struct {
	ChatID   uuid.UUID
	Display  string
	Activity string
}
