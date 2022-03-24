package repositories

import (
	"api/src/model"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create a new post
func (repo Posts) Create(post model.Post) (uint64, error) {
	statement, err := repo.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}
