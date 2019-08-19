package mutations

import (
	"github.com/disturb16/graphql_golang/internal/services"
	"github.com/disturb16/graphql_golang/schema/types"
	"github.com/graphql-go/graphql"
)

// AddCommentMutation saves comment in database
var AddCommentMutation = &graphql.Field{
	Type: types.CommentType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"content": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"postId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		service := params.Context.Value("service").(*services.Service)

		name := params.Args["name"].(string)
		content := params.Args["content"].(string)
		postID := params.Args["postId"].(int)

		insertID, err := service.AddComment(name, content, postID)

		if err != nil {
			return nil, err
		}

		comment, err := service.CommentByID(insertID)

		if err != nil {
			return nil, err
		}

		return comment, nil
	},
}
