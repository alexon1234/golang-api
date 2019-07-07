package domain

// Post - Domain of the post
type Post struct {
	ID    string `json:"ID"`
	Title string `json:"Title"`
	Body  string `json:"Body"`
}

// Repository - interface of the Post Repository
type Repository interface {
	RunMigrations()
	FetchPosts() (*[]Post, error)
	GetPost(id string) (*Post, error)
	CreatePost(post *Post) error
}
