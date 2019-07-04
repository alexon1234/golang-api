package storage

import (
	post "github.com/alexon1234/golang-api/pkg"
	"github.com/go-pg/pg"
)

type postRepository struct {
}

// NewPostRepository - Create a new Post Repository
func NewPostRepository() post.Repository {
	return &postRepository{}
}

func (r *postRepository) FetchPosts() (*[]post.Post, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "password",
	})
	defer db.Close()
	var posts []post.Post
	if err := db.Model(&posts).Select(); err != nil {
		return nil, err
	}

	return &posts, nil
}
