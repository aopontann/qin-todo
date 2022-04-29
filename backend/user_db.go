package main

import (
	"database/sql"
)

func GetUser(userId string) (GetUserResponse, error) {
	var (
		id         string
		name       string
		email      string
		avatar_url *sql.NullString
	)
	err := db.QueryRow("SELECT id, name, email, avatar_url FROM users WHERE id = ?", userId).Scan(&id, &name, &email, &avatar_url)
	if err != nil {
		return GetUserResponse{}, err
	}

	if avatar_url != nil {
		return GetUserResponse{id, name, email, avatar_url.String}, nil
	}
	return GetUserResponse{id, name, email, ""}, nil
}

func DeleteUser(userId string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM todo_list WHERE user_id = ?", userId)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM users WHERE id = ? LIMIT 1", userId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
