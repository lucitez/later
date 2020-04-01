package repository

import (
	"database/sql"

	"later/pkg/model"
	"later/pkg/repository/util"

	"github.com/google/uuid"

	// Postgres driver
	_ "github.com/lib/pq"
)

// Share ...
type Share struct {
	DB *sql.DB
}

// NewShare ...
func NewShare(db *sql.DB) Share {
	return Share{db}
}

var selectShares = util.GenerateSelectStatement(model.Share{}, "shares")

// Insert inserts a new share
func (repository *Share) Insert(share *model.Share) (*model.Share, error) {

	statement := util.GenerateInsertStatement(*share, "shares")

	_, err := repository.DB.Exec(
		statement,
		share.ID,
		share.ContentID,
		share.SentByUserID,
		share.RecipientUserID,
		share.CreatedAt,
		share.OpenedAt)

	if err != nil {
		return nil, err
	}

	return share, nil
}

// ByID gets a share by id
func (repository *Share) ByID(id uuid.UUID) (*model.Share, error) {
	var share model.Share

	statement := selectShares + `
	WHERE id = $1
	`

	row := repository.DB.QueryRow(statement, id)

	err := share.ScanRow(row)

	return &share, err
}

// All returns all shares
func (repository *Share) All(limit int) ([]model.Share, error) {
	statement := selectShares + `
	LIMIT $1;
	`

	rows, err := repository.DB.Query(statement, limit)

	if err != nil {
		return nil, err
	}

	return repository.scanRows(rows)
}

func (repository *Share) scanRows(rows *sql.Rows) ([]model.Share, error) {
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
