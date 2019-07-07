package infrastructure

import (
	"fmt"

	"github.com/alexon1234/golang-api/pkg/post/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type postRepository struct {
	*gorm.DB
}

// NewPostRepository - Create a new Post Repository
func NewPostRepository() domain.Repository {
	db, err := gorm.Open("postgres", "host=192.168.99.101 port=5432 user=root dbname=api sslmode=disable password=root")

	if err != nil {
		panic("Fail to connect database")
	}

	return &postRepository{db}
}

func (r *postRepository) RunMigrations() {
	fmt.Println("Running Migrations")
	r.AutoMigrate(domain.Post{})
}

func (r *postRepository) GetPost(id string) (*domain.Post, error) {
	var post domain.Post
	err := r.First(&post, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &post, err
}

func (r *postRepository) FetchPosts() (*[]domain.Post, error) {

	var posts []domain.Post
	if err := r.Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}

func (r *postRepository) CreatePost(post *domain.Post) (err error) {
	err = r.Create(post).Error
	return
}
