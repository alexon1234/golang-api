package main

import (
	"net/http"

	"github.com/alexon1234/golang-api/pkg/server"
	"github.com/alexon1234/golang-api/pkg/storage"
)

func main() {

	repo := storage.NewPostRepository()
	s := server.New(repo)

	http.ListenAndServe(":8000", s.Router())
}
