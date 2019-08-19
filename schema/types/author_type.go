package types

import (
	"github.com/disturb16/graphql_golang/internal/models"
	"github.com/disturb16/graphql_golang/internal/services"
	"github.com/graphql-go/graphql"
)

// AuthorType graphqlObject for authors
var AuthorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"posts": &graphql.Field{
			Type: graphql.NewList(PostType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				service := params.Context.Value("service").(*services.Service)

				author := params.Source.(models.Author)
				result, err := service.PostsByAuthor(author.ID)
				if err != nil {
					return nil, err
				}
				return result, nil
			},
		},
	},
})
