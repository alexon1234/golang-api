package main

import (
	"net/http"

	"github.com/alexon1234/golang-api/pkg"
)

func main() {
	s := pkg.New()
	http.ListenAndServe(":8000", s.Router())
}
