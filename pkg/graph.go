package pkg

import (
	"fmt"

	"github.com/alexon1234/golang-api/pkg/post/infrastructure"
	"github.com/graphql-go/graphql"
)

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:   "RootMutation",
	Fields: graphql.Fields{},
})

var postsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"body": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// root query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"posts": &graphql.Field{
			Type: graphql.NewList(postsType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				repo := infrastructure.NewPostRepository()
				posts, err := repo.FetchPosts()
				return *posts, err
			},
		},
	},
})

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

// ExecuteQuery -
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func NewSchema() *graphql.Schema {
	return &schema
}
