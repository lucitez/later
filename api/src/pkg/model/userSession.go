package model

import (
	"database/sql"
	"later/pkg/errs"
	"later/pkg/util/wrappers"
	"time"

	"github.com/google/uuid"
)

// UserSession contains information about a users current userSession
type UserSession struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	ExpiresAt time.Time
	ExpiredAt wrappers.NullTime
}

// NewUserSession constructor for a UserSession
func NewUserSession(UserID uuid.UUID) UserSession {
	now := time.Now()
	id, _ := uuid.NewRandom()

	return UserSession{
		ID:        id,
		UserID:    UserID,
		CreatedAt: now,
		ExpiresAt: now.Add(time.Minute * 15),
	}
}

// ScanRow ...
func (userSession *UserSession) ScanRow(row *sql.Row) (*UserSession, error) {
	err := row.Scan(
		&userSession.ID,
		&userSession.UserID,
		&userSession.CreatedAt,
		&userSession.ExpiresAt,
		&userSession.ExpiredAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &errs.QueryRowUnknown{Tablename: "user_sessions", Err: err}
	}

	return userSession, nil
}
