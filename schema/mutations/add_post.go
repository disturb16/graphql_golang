package mutations

import (
	"github.com/disturb16/graphql_golang/internal/services"
	"github.com/disturb16/graphql_golang/schema/types"
	"github.com/graphql-go/graphql"
)

// AddPostMutation saves post in database
var AddPostMutation = &graphql.Field{
	Type: types.PostType,
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"content": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"authorId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		service := params.Context.Value("service").(*services.Service)

		title := params.Args["title"].(string)
		content := params.Args["content"].(string)
		authorID := params.Args["authorId"].(int)

		insertID, err := service.AddPost(title, content, authorID)

		if err != nil {
			return nil, err
		}

		post, err := service.PostByID(insertID)

		if err != nil {
			return nil, err
		}

		return post, nil
	},
}
