package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
)

type PostHandler struct {
}

func NewPostHandler() *PostHandler {
	return &PostHandler{}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
}
