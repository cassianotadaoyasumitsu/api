package repositories

import (
	"api/src/model"
	"database/sql"
	"fmt"
)

// Representa repo
type Users struct {
	db *sql.DB
}

// Cria um repo de usuários
func NewUsersRepo(db *sql.DB) *Users {
	return &Users{db}
}

// Cria usuário no DB
func (repo Users) Create(user model.User) (uint64, error) {
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
func (repo Users) Search(nameOrNick string) ([]model.User, error) {
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
func (repo Users) SearchByID(ID uint64) (model.User, error) {
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

// User update
func (repo Users) Update(ID uint64, user model.User) error {
	statement, err := repo.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// User delete
func (repo Users) Delete(ID uint64) error {
	statement, err := repo.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// User search by email
func (repo Users) SearchByEmail(email string) (model.User, error) {
	line, err := repo.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return model.User{}, err
	}
	defer line.Close()

	var user model.User

	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

// Follow user
func (repo Users) Follow(followerID, userID uint64) error {
	statement, err := repo.db.Prepare("insert ignore into followers (follower_id, user_id) values(?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(followerID, userID); err != nil {
		return err
	}

	return nil
}

// Unfollow user
func (repo Users) Unfollow(followerID, userID uint64) error {
	statement, err := repo.db.Prepare("delete from followers where follower_id = ? and user_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(followerID, userID); err != nil {
		return err
	}

	return nil
}

// Get followers
func (repo Users) GetFollowers(userID uint64) ([]model.User, error) {
	lines, err := repo.db.Query("select id, name, nick, email, created_at from users where id in (select follower_id from followers where user_id = ?)", userID)
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

// Get following
func (repo Users) GetFollowing(userID uint64) ([]model.User, error) {
	lines, err := repo.db.Query("select id, name, nick, email, created_at from users where id in (select user_id from followers where follower_id = ?)", userID)
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

// Get password
func (repo Users) GetPassword(userID uint64) (string, error) {
	line, err := repo.db.Query("select password from users where id = ?", userID)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user model.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

// Password update
func (repo Users) UpdatePassword(userID uint64, password string) error {
	statement, err := repo.db.Prepare("update users set password = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userID); err != nil {
		return err
	}

	return nil
}
