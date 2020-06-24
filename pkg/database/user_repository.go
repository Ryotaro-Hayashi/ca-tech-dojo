package database

import (
	"ca-tech-dojo/internal/game_api/user/models"
)

type UserRepository struct {
	*SqlHandler
}

// DBからユーザー一覧を取得して返す
func (repo UserRepository) GetUsers() (users models.Users, err error){
	rows, err := repo.SqlHandler.Conn.Query("SELECT name, token FROM users")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var token string

		if err = rows.Scan(&name, &token); err != nil {
			return
		}

		user := models.User {
			Name: name,
			Token: token,
		}

		users = append(users, user)
	}

	return
}