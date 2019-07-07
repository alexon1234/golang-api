package pkg

import (
	"net/http"

	"github.com/alexon1234/golang-api/pkg/post/infrastructure"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

// Server - interface use to initialize the api
type Server interface {
	Router() http.Handler
}

// New - To create a new Server
func New() Server {
	repo := infrastructure.NewPostRepository()
	postController := infrastructure.NewPostController(repo)

	r := mux.NewRouter()
	r.HandleFunc("/posts", postController.FetchPosts).Methods(http.MethodGet)
	r.HandleFunc("/posts/{id}", postController.GetPost).Methods(http.MethodGet)
	r.HandleFunc("/posts", postController.CreatePost).Methods(http.MethodPost)

	return &api{router: r}
}

func (a *api) Router() http.Handler {
	return a.router
}
