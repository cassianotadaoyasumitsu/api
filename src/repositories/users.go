package repositories

import (
	"api/src/model"
	"database/sql"
	"fmt"
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

// Busca usuários no DB por filtro name nick
func (repo users) Search(nameOrNick string) ([]model.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, err := repo.db.Query(
		"select id, name, nick, email, created_at from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []model.User

	for lines.Next() {
		var user model.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Created_at,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Busca usuário por ID
func (repo users) SearchByID(ID uint64) (model.User, error) {
	lines, err := repo.db.Query(
		"select id, name, nick, email, created_at from users where id = ?",
		ID,
	)
	if err != nil {
		return model.User{}, err
	}
	defer lines.Close()

	var user model.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Created_at,
		); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}
