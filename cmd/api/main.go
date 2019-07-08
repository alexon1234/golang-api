package main

import (
	"net/http"

	"github.com/alexon1234/golang-api/pkg"
	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/handler"
)

func main() {
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}
	h := handler.New(&handler.Config{
		Schema:   pkg.NewSchema(),
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	http.Handle("/graphiql", graphiqlHandler)
	http.ListenAndServe(":8000", nil)
}
