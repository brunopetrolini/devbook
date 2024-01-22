package repositories

import "database/sql"

type followers struct {
	db *sql.DB
}

func FollowersRepository(db *sql.DB) *followers {
	return &followers{db}
}

// FollowUser follows a user
func (f followers) FollowUser(userID, followerID uint64) error {
	statement, error := f.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(userID, followerID); error != nil {
		return error
	}

	return nil
}

// UnfollowUser unfollows a user
func (f followers) UnfollowUser(userID, followerID uint64) error {
	statement, error := f.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(userID, followerID); error != error {
		return error
	}

	return nil
}
