package database

import (
	"ca-tech-dojo/internal/game_api/user/models"
)

type UserRepository struct {
	*SqlHandler
}

// DBからユーザー一覧を取得して返す
func (repo *UserRepository) GetAll() (users models.Users, err error){
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

// DBにユーザーを保存して、保存したユーザーidを返す
func (repo *UserRepository) Create(name string, tokenString string) (id int64, err error) {
	result, err := repo.SqlHandler.Conn.Exec("INSERT INTO users (name, token) VALUES (?, ?)", name, tokenString)
	if err != nil {
		return
	}

	id, err = result.LastInsertId()
	if err != nil {
		return
	}

	return
}

// DBに保存されているユーザーのトークンをidで検索して返す
func (repo *UserRepository) FindTokenById(id int64) (tokenString string, err error) {
	row, err := repo.SqlHandler.Conn.Query("SELECT token FROM users where id = ?", id)
	if err != nil {
		return
	}
	defer row.Close()

	row.Next()
	if err = row.Scan(&tokenString); err != nil {
		return
	}
	
	return
}