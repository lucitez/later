package model

import (
	"database/sql"
	"later/pkg/errs"
	"later/pkg/util/wrappers"
	"time"

	"github.com/google/uuid"
)

// Session contains information about a users current session
type Session struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	ExpiresAt time.Time
	ExpiredAt wrappers.NullTime
}

// NewSession constructor for a Session
func NewSession(UserID uuid.UUID) Session {
	now := time.Now()
	id, _ := uuid.NewRandom()

	return Session{
		ID:        id,
		UserID:    UserID,
		CreatedAt: now,
		ExpiresAt: now.Add(time.Minute * 15),
	}
}

// ScanRow ...
func (session *Session) ScanRow(row *sql.Row) (*Session, error) {
	err := row.Scan(
		&session.ID,
		&session.UserID,
		&session.CreatedAt,
		&session.ExpiresAt,
		&session.ExpiredAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &errs.QueryRowUnknown{Tablename: "sessions", Err: err}
	}

	return session, nil
}
