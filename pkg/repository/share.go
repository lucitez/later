package repository

import (
	"database/sql"
	"log"

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

var shareSelectStatement = util.GenerateSelectStatement(model.Share{}, "shares")

// Insert inserts a new share
func (repository *Share) Insert(share model.Share) error {
	statement := util.GenerateInsertStatement(share, "shares")

	_, err := repository.DB.Exec(
		statement,
		share.ID,
		share.ContentID,
		share.SentByUserID,
		share.RecipientUserID,
		share.CreatedAt,
		share.OpenedAt,
	)

	return err
}

// ByID gets a share by id
func (repository *Share) ByID(id uuid.UUID) *model.Share {
	var share model.Share

	statement := shareSelectStatement + `
	WHERE id = $1;
	`

	row := repository.DB.QueryRow(statement, id)

	share.ScanRow(row)

	return &share
}

// All returns all shares
func (repository *Share) All(limit int) []model.Share {
	statement := shareSelectStatement + `
	LIMIT $1;
	`

	rows, err := repository.DB.Query(statement, limit)

	if err != nil {
		log.Fatal(err)
	}

	return repository.scanRows(rows)
}

func (repository *Share) scanRows(rows *sql.Rows) []model.Share {
	shares := []model.Share{}

	defer rows.Close()

	for rows.Next() {
		var share model.Share
		share.ScanRows(rows)

		shares = append(shares, share)
	}

	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return shares
}
