package storage

import (
	"fmt"

	post "github.com/alexon1234/golang-api/pkg"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type postRepository struct {
	*gorm.DB
}

// NewPostRepository - Create a new Post Repository
func NewPostRepository() post.Repository {
	db, err := gorm.Open("postgres", "host=192.168.99.101 port=5432 user=root dbname=api sslmode=disable password=root")

	if err != nil {
		panic("Fail to connect database")
	}

	return &postRepository{db}
}

func (r *postRepository) RunMigrations() {
	fmt.Println("Running Migrations")
	r.AutoMigrate(post.Post{})
}

func (r *postRepository) FetchPosts() (*[]post.Post, error) {

	var posts []post.Post
	if err := r.Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}

func (r *postRepository) CreatePost(post *post.Post) (err error) {
	err = r.Create(post).Error
	return
}
