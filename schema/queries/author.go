package queries

import (
	"github.com/disturb16/graphql_golang/internal/services"
	"github.com/disturb16/graphql_golang/schema/types"
	"github.com/graphql-go/graphql"
)

// AuhtorQuery returns specific author
var AuhtorQuery = &graphql.Field{
	Type: types.AuthorType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		service := params.Context.Value("service").(*services.Service)

		author, err := service.GetAuthorByID(params.Args["id"].(int))

		if err != nil {
			return nil, err
		}

		return author, nil
	},
}
