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

// Search posts by user id
func (repo Posts) SearchByID(postID uint64) (model.Post, error) {
	line, err := repo.db.Query("select p.*, u.nick from posts p inner join users u on u.id = p.author_id where p.id = ?", postID)
	if err != nil {
		return model.Post{}, err
	}
	defer line.Close()

	var post model.Post

	if line.Next() {
		if err = line.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.Created_at,
			&post.AuthorNick,
		); err != nil {
			return model.Post{}, err
		}
	}

	return post, nil
}

// Search posts
func (repo Posts) Search(userID uint64) ([]model.Post, error) {
	lines, err := repo.db.Query("select distinct p.*, u.nick from posts p inner join users u on u.id = p.author_id inner join followers s on p.author_id = s.user_id where u.id = ? or s.follower_id = ? order by 1 desc", userID, userID)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []model.Post

	for lines.Next() {
		var post model.Post
		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.Created_at,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Update post
func (repo Posts) Update(postID uint64, post model.Post) error {
	statement, err := repo.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

// Delete post
func (repo Posts) Delete(postID uint64) error {
	statement, err := repo.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repo Posts) SearchByUserID(userID uint64) ([]model.Post, error) {
	lines, err := repo.db.Query("select p.*, u.nick from posts p inner join users u on u.id = p.author_id where p.author_id = ? order by 1 desc", userID)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []model.Post

	for lines.Next() {
		var post model.Post
		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.Created_at,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
