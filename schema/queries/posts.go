package queries

import (
	"github.com/disturb16/graphql_golang/internal/services"
	"github.com/disturb16/graphql_golang/schema/types"
	"github.com/graphql-go/graphql"
)

// PostsQuery returns the list of posts
var PostsQuery = &graphql.Field{
	Type: graphql.NewList(types.PostType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		service := params.Context.Value("service").(*services.Service)

		posts, err := service.GetPosts()
		if err != nil {
			return nil, err
		}

		return posts, nil
	},
}
