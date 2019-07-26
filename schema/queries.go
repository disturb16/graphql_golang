package schema

import (
	"github.com/graphql-go/graphql"
)

// postsQuery returns the list of posts
var postsQuery = &graphql.Field{
	Type: graphql.NewList(PostType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		posts, err := service.GetPosts()

		if err != nil {
			return nil, err
		}

		return posts, nil
	},
}

var postQuery = &graphql.Field{
	Type: PostType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Args["id"].(int)
		post, err := service.GetPostByID(int64(id))

		if err != nil {
			return nil, err
		}

		return post, nil
	},
}

var auhtorQuery = &graphql.Field{
	Type: AuthorType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		author, err := service.GetAuthorByID(params.Args["id"].(int))

		if err != nil {
			return nil, err
		}

		return author, nil
	},
}
