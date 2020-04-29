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

// InsertUserSession inserts a new userSession
func (repo *Auth) InsertUserSession(userSession model.UserSession) error {
	statement := util.GenerateInsertStatement(userSession, "user_sessions")

	_, err := repo.DB.Exec(
		statement,
		userSession.ID,
		userSession.UserID,
		userSession.CreatedAt,
		userSession.ExpiresAt,
		userSession.ExpiredAt,
	)

	return err
}

// ExpireUserSession expires a user userSession
func (repo *Auth) ExpireUserSession(id uuid.UUID) {
	statement := `
	UPDATE user_sessions
	SET expired_at = now()
	WHERE id = $1;
	`

	if _, err := repo.DB.Exec(statement, id); err != nil {
		log.Fatal(err)
	}
}

// ByID get a userSession
func (repo *Auth) ByID(id uuid.UUID) (*model.UserSession, error) {
	userSession := model.UserSession{}

	statement := `
	SELECT * FROM user_sessions
	WHERE ID = $1;
	`

	row := repo.DB.QueryRow(statement, id)

	return userSession.ScanRow(row)
}

// ActiveByID get a userSession
func (repo *Auth) ActiveByID(id uuid.UUID) (*model.UserSession, error) {
	userSession := model.UserSession{}

	statement := `
	SELECT * FROM user_sessions
	WHERE ID = $1
	AND expires_at > now();
	`

	row := repo.DB.QueryRow(statement, id)

	return userSession.ScanRow(row)
}
