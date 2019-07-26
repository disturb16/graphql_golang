package schema

import (
	"github.com/disturb16/graphql_golang/internal/models"
	"github.com/graphql-go/graphql"
)

// PostType graphqlObject for posts
var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"comments": &graphql.Field{
			Type: graphql.NewList(CommentType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				post := params.Source.(models.Post)
				comments, err := service.GetCommentsByPost(post.ID)

				if err != nil {
					return nil, err
				}

				return comments, nil
			},
		},
	},
})
