package repository

import (
	"database/sql"

	"later/pkg/model"

	"github.com/google/uuid"

	// Postgres driver
	_ "github.com/lib/pq"
)

// FriendRepository ...
type FriendRepository struct {
	DB *sql.DB
}

// NewFriendRepository for wire generation
func NewFriendRepository(db *sql.DB) FriendRepository {
	return FriendRepository{db}
}

// Insert inserts a new friend
func (repository *FriendRepository) Insert(friend *model.Friend) (*model.Friend, error) {

	statement := `
	INSERT INTO friends (id, user_id, friend_user_id)
	VALUES (
		$1,
		$2,
		$3
	)
	`

	_, err := repository.DB.Exec(
		statement,
		friend.ID,
		friend.UserID,
		friend.FriendUserID)

	if err != nil {
		return nil, err
	}

	return friend, nil
}

// SearchByUserID gets all friends of a user
func (repository *FriendRepository) SearchByUserID(userID uuid.UUID, search string) ([]model.Friend, error) {
	statement := `
	SELECT * FROM friends 
	WHERE user_id = $1
	AND (
		OR (
			users.username ilike %$2%,
			users.email ilike %$2%,
			users.first_name ilike %$2%,
			users.last_name ilike %$2%
		)
	)
	AND deleted_at IS NULL
	`

	rows, err := repository.DB.Query(statement, userID, search)

	if err != nil {
		return nil, err
	}

	friends, err := repository.scanRows(rows)

	if err != nil {
		return nil, err
	}

	return friends, nil
}

// ByUserID gets all friends of a user
func (repository *FriendRepository) ByUserID(userID uuid.UUID) ([]model.Friend, error) {
	statement := `
	SELECT * FROM friends 
	WHERE user_id = $1
	AND deleted_at IS NULL
	`

	rows, err := repository.DB.Query(statement, userID)

	if err != nil {
		return nil, err
	}

	friends, err := repository.scanRows(rows)

	if err != nil {
		return nil, err
	}

	return friends, nil
}

func (repository *FriendRepository) scanRows(rows *sql.Rows) ([]model.Friend, error) {
	friends := []model.Friend{}

	defer rows.Close()

	for rows.Next() {
		var friend model.Friend
		err := friend.ScanRows(rows)

		if err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}

	err := rows.Err()
	if err != nil {
		return nil, err
	}

	return friends, nil
}
