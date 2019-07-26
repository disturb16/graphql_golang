package queries

import (
	"github.com/disturb16/graphql_golang/internal/services"
	"github.com/disturb16/graphql_golang/schema/types"
	"github.com/graphql-go/graphql"
)

// PostQuery returns specific post
var PostQuery = &graphql.Field{
	Type: types.PostType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		service := params.Context.Value("service").(*services.Service)

		id := params.Args["id"].(int)
		post, err := service.GetPostByID(int64(id))

		if err != nil {
			return nil, err
		}

		return post, nil
	},
}
