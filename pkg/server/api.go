package server

import (
	"encoding/json"
	"net/http"

	post "github.com/alexon1234/golang-api/pkg"
	"github.com/gorilla/mux"
)

type api struct {
	router     http.Handler
	repository post.Repository
}

// Server - interface use to initialize the api
type Server interface {
	Router() http.Handler
}

// New - To create a new Server
func New(repo post.Repository) Server {
	a := &api{repository: repo}

	r := mux.NewRouter()
	r.HandleFunc("/posts", a.fetchPosts).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchPosts(w http.ResponseWriter, r *http.Request) {
	gophers, _ := a.repository.FetchPosts()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gophers)
}
