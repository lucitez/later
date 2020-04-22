package repository

import (
	"database/sql"
	"later/pkg/model"
	"later/pkg/repository/util"
	"log"

	"github.com/google/uuid"
)

// Auth ...
type Auth struct {
	DB *sql.DB
}

// NewAuth for wire generation
func NewAuth(db *sql.DB) Auth {
	return Auth{db}
}

// InsertSession inserts a new session
func (repo *Auth) InsertSession(session model.Session) error {
	statement := util.GenerateInsertStatement(session, "sessions")

	_, err := repo.DB.Exec(
		statement,
		session.ID,
		session.UserID,
		session.CreatedAt,
		session.ExpiresAt,
		session.ExpiredAt,
	)

	return err
}

// ExpireSession expires a user session
func (repo *Auth) ExpireSession(id uuid.UUID) {
	statement := `
	UPDATE sessions
	SET expired_at = now()
	WHERE id = $1;
	`

	if _, err := repo.DB.Exec(statement, id); err != nil {
		log.Fatal(err)
	}
}

// ByID get a session
func (repo *Auth) ByID(id uuid.UUID) (*model.Session, error) {
	session := model.Session{}

	statement := `
	SELECT * FROM sessions
	WHERE ID = $1;
	`

	row := repo.DB.QueryRow(statement, id)

	return session.ScanRow(row)
}

// ActiveByID get a session
func (repo *Auth) ActiveByID(id uuid.UUID) (*model.Session, error) {
	session := model.Session{}

	statement := `
	SELECT * FROM sessions
	WHERE ID = $1
	AND expires_at > now();
	`

	row := repo.DB.QueryRow(statement, id)

	return session.ScanRow(row)
}
