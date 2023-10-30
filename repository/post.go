package repository

import "time"

type Post struct {
	ID        int64     `db:"id" validate:"required"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
}

func (r *Repository) CreatePost(title string, body string) (int64, error) {
	res, err := r.DB.Exec("INSERT INTO posts (title, body) VALUES ($1, $2)", title, body)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *Repository) GetPosts() (*[]Post, error) {
	posts := []Post{}

	err := r.DB.Select(&posts, "SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	return &posts, nil
}
