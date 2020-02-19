package later

import (
	"database/sql"

	"github.com/google/uuid"

	"time"
)

// EntityUser defines the user
type EntityUser struct {
	ID          uuid.UUID
	Username    string
	Email       sql.NullString
	Phonenumber sql.NullInt32

	CreatedAt  time.Time
	SignedUpAt sql.NullTime
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime
}

// UserSignUpRequestBody Binding from json
type UserSignUpRequestBody struct {
	Username    string `form:"user_name" json:"user_name" binding:"required"`
	Email       sql.NullString `form:"email" json:"email"`
	Phonenumber sql.NullInt32 `form:"phonenumber" json:"phonenumber"`
}
