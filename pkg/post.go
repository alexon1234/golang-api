package post

// Post - Domain of the post
type Post struct {
	ID    string `json:"ID"`
	Title string `json:"Title"`
	Body  string `json:"Body"`
}

// Repository - interface of the Post Repository
type Repository interface {
	FetchPosts() (*[]Post, error)
	RunMigrations()
	CreatePost(post *Post) error
}
