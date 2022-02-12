package repositories

import (
	"api/src/model"
	"database/sql"
)

// Representa repo
type users struct {
	db *sql.DB
}

// Cria um repo de usuários
func NewUsersRepo(db *sql.DB) *users {
	return &users{db}
}

// Cria usuário no DB
func (repo users) Create(user model.User) (uint64, error) {
	statement, err := repo.db.Prepare("insert into users (name, nick, email, password) values(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil
}
