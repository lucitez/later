package friendrepo

import (
	"database/sql"

	// Postgres driver
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"later.co/pkg/later/friend"
)

// DB is this repository's database connection
var DB *sql.DB

// Insert inserts a new friend
func Insert(newFriend *friend.Friend) (*friend.Friend, error) {

	statement := `
	INSERT INTO friends (id, user_id, friend_user_id)
	VALUES (
		$1,
		$2,
		$3
	)
	`

	_, err := DB.Exec(
		statement,
		newFriend.ID,
		newFriend.UserID,
		newFriend.FriendUserID)

	if err != nil {
		return nil, err
	}

	return newFriend, nil
}

// SearchByUserID gets all friends of a user
func SearchByUserID(userID uuid.UUID, search string) ([]friend.Friend, error) {
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

	rows, err := DB.Query(statement, userID, search)

	if err != nil {
		return nil, err
	}

	friends, err := scanRows(rows)

	if err != nil {
		return nil, err
	}

	return friends, nil
}

// ByUserID gets all friends of a user
func ByUserID(userID uuid.UUID) ([]friend.Friend, error) {
	statement := `
	SELECT * FROM friends 
	WHERE user_id = $1
	AND deleted_at IS NULL
	`

	rows, err := DB.Query(statement, userID)

	if err != nil {
		return nil, err
	}

	friends, err := scanRows(rows)

	if err != nil {
		return nil, err
	}

	return friends, nil
}

func scanRows(rows *sql.Rows) ([]friend.Friend, error) {
	friends := []friend.Friend{}

	defer rows.Close()

	for rows.Next() {
		var friend friend.Friend
		err := rows.Scan(
			&friend.ID,
			&friend.UserID,
			&friend.FriendUserID,
			&friend.CreatedAt,
			&friend.UpdatedAt,
			&friend.DeletedAt)

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
