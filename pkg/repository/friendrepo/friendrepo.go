package friendrepo

import (
	"database/sql"

	"github.com/google/uuid"
	// Postgres driver
	_ "github.com/lib/pq"
	"later.co/pkg/later/entity"
)

// DB is this repository's database connection
var DB *sql.DB

// Insert inserts a new friend
func Insert(friend *entity.Friend) (*entity.Friend, error) {

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
		friend.ID,
		friend.UserID,
		friend.FriendUserID)

	if err != nil {
		return nil, err
	}

	return friend, nil
}

// SearchByUserID gets all friends of a user
func SearchByUserID(userID uuid.UUID, search string) ([]entity.Friend, error) {
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
func ByUserID(userID uuid.UUID) ([]entity.Friend, error) {
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

func scanRows(rows *sql.Rows) ([]entity.Friend, error) {
	friends := []entity.Friend{}

	defer rows.Close()

	for rows.Next() {
		var friend entity.Friend
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
