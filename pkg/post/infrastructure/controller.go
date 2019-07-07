package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alexon1234/golang-api/pkg/post/domain"
	"github.com/gorilla/mux"
)

// PostController -
type PostController struct {
	repository domain.Repository
}

// NewPostController - New Post Controller
func NewPostController(repo domain.Repository) *PostController {
	return &PostController{repository: repo}
}

func (p *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post domain.Post
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.Unmarshal(data, &post)
	if err := p.repository.CreatePost(&post); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (p *PostController) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	post, err := p.repository.GetPost(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (p *PostController) FetchPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := p.repository.FetchPosts()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
