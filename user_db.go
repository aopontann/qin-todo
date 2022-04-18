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
