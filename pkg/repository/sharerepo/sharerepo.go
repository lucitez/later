package sharerepo

import (
	"database/sql"

	// Postgres driver
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"later.co/pkg/later/share"
)

// DB is this repository's database connection
var DB *sql.DB

// Insert inserts a new share
func Insert(share *share.Share) (*share.Share, error) {

	statement := `
	INSERT INTO shares (id, content_id, sent_by_user_id, recipient_user_id)
	VALUES (
		$1,
		$2,
		$3,
		$4
	)
	`

	_, err := DB.Exec(
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
func ByID(id uuid.UUID) (*share.Share, error) {
	var share share.Share

	statement := `
	SELECT * FROM shares 
	WHERE id = $1
	`

	row := DB.QueryRow(statement, id)

	err := row.Scan(
		&share.ID,
		&share.ContentID,
		&share.SentByUserID,
		&share.RecipientUserID,
		&share.CreatedAt,
		&share.OpenedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &share, nil
}

// All returns all shares
func All(limit int) ([]share.Share, error) {
	shares := []share.Share{}

	rows, err := DB.Query(`SELECT * FROM shares LIMIT $1`, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var share share.Share
		err := rows.Scan(
			&share.ID,
			&share.ContentID,
			&share.SentByUserID,
			&share.RecipientUserID,
			&share.CreatedAt,
			&share.OpenedAt)

		if err != nil {
			return nil, err
		}
		shares = append(shares, share)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return shares, nil
}
