package repository

import (
	"database/sql"

	"later/pkg/model"
	"later/pkg/repository/util"

	"github.com/google/uuid"

	// Postgres driver
	_ "github.com/lib/pq"
)

// Friend ...
type Friend struct {
	DB *sql.DB
}

// NewFriend for wire generation
func NewFriend(db *sql.DB) Friend {
	return Friend{db}
}

var friendSelectStatement = util.GenerateSelectStatement(model.Friend{}, "friends")

// Insert inserts a new friend
func (repository *Friend) Insert(friend *model.Friend) (*model.Friend, error) {

	statement := util.GenerateInsertStatement(*friend, "friends")

	_, err := repository.DB.Exec(
		statement,
		friend.ID,
		friend.UserID,
		friend.FriendUserID,
		friend.CreatedAt,
		friend.UpdatedAt,
		friend.DeletedAt)

	if err != nil {
		return nil, err
	}

	return friend, nil
}

// SearchByUserID gets all friends of a user
func (repository *Friend) SearchByUserID(userID uuid.UUID, search string) ([]model.Friend, error) {
	search = "%" + search + "%"
	statement := friendSelectStatement + `
	JOIN users on users.id = friends.friend_user_id
	WHERE friends.user_id = $1
	AND (
		users.username ILIKE $2
		OR users.email ILIKE $2
		OR users.first_name ILIKE $2
		OR users.last_name ILIKE $2
	)
	AND friends.deleted_at IS NULL;
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
func (repository *Friend) ByUserID(userID uuid.UUID) ([]model.Friend, error) {
	statement := friendSelectStatement + `
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

func (repository *Friend) scanRows(rows *sql.Rows) ([]model.Friend, error) {
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
