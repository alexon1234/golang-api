package main

import (
	"net/http"

	api "github.com/golang-api/internal"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/posts", api.NewPostHandler())
	http.ListenAndServe(":8000", r)
}
