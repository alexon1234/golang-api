package domain

// Post - Domain of the post
type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Repository - interface of the Post Repository
type Repository interface {
	RunMigrations()
	FetchPosts() (*[]Post, error)
	GetPost(id string) (*Post, error)
	CreatePost(post *Post) error
}
