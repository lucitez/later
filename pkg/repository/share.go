package repository

import (
	"database/sql"

	"github.com/google/uuid"
	// Postgres driver
	_ "github.com/lib/pq"
	"later/pkg/model"
)

// ShareRepository ...
type ShareRepository struct {
	DB *sql.DB
}

// NewShareRepository ...
func NewShareRepository(db *sql.DB) ShareRepository {
	return ShareRepository{db}
}

// Insert inserts a new share
func (repository *ShareRepository) Insert(share *model.Share) (*model.Share, error) {

	statement := `
	INSERT INTO shares (id, content_id, sent_by_user_id, recipient_user_id)
	VALUES (
		$1,
		$2,
		$3,
		$4
	)
	`

	_, err := repository.DB.Exec(
		statement,
		share.ID,
		share.ContentID,
		share.SentByUserID,
		share.RecipientUserID)

	if err != nil {
		return nil, err
	}

	return share, nil
}

// ByID gets a share by id
func (repository *ShareRepository) ByID(id uuid.UUID) (*model.Share, error) {
	var share model.Share

	statement := `
	SELECT * FROM shares 
	WHERE id = $1
	`

	row := repository.DB.QueryRow(statement, id)

	err := share.ScanRow(row)

	return &share, err
}

// All returns all shares
func (repository *ShareRepository) All(limit int) ([]model.Share, error) {
	statement := `
	SELECT * FROM shares
	WHERE deleted_at IS NULL
	LIMIT $1;
	`

	rows, err := repository.DB.Query(statement, limit)

	if err != nil {
		return nil, err
	}

	return repository.scanRows(rows)
}

func (repository *ShareRepository) scanRows(rows *sql.Rows) ([]model.Share, error) {
	shares := []model.Share{}

	defer rows.Close()

	for rows.Next() {
		var share model.Share
		err := share.ScanRows(rows)

		if err != nil {
			return nil, err
		}
		shares = append(shares, share)
	}

	err := rows.Err()
	if err != nil {
		return nil, err
	}

	return shares, nil
}
