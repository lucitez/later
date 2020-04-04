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
func (repository *Friend) Insert(friend model.Friend) error {

	statement := util.GenerateInsertStatement(friend, "friends")

	_, err := repository.DB.Exec(
		statement,
		friend.ID,
		friend.UserID,
		friend.FriendUserID,
		friend.CreatedAt,
		friend.UpdatedAt,
		friend.DeletedAt,
	)

	return err
}

// SearchByUserID gets all friends of a user
func (repository *Friend) SearchByUserID(userID uuid.UUID, search string) []model.Friend {
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
		log.Fatal(err)
	}

	return repository.scanRows(rows)
}

// ByUserID gets all friends of a user
func (repository *Friend) ByUserID(userID uuid.UUID) []model.Friend {
	statement := friendSelectStatement + `
	WHERE user_id = $1
	AND deleted_at IS NULL
	`

	rows, err := repository.DB.Query(statement, userID)

	if err != nil {
		log.Fatal(err)
	}

	friends := repository.scanRows(rows)

	return friends
}

func (repository *Friend) scanRows(rows *sql.Rows) []model.Friend {
	friends := []model.Friend{}

	defer rows.Close()

	for rows.Next() {
		var friend model.Friend
		friend.ScanRows(rows)

		friends = append(friends, friend)
	}

	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return friends
}
