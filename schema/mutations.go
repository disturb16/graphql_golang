package schema

import "github.com/graphql-go/graphql"

var addPostMutation = &graphql.Field{
	Type: PostType,
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

		title := params.Args["title"].(string)
		content := params.Args["content"].(string)
		authorId := params.Args["authorId"].(int)
		insertId, err := service.AddPost(title, content, authorId)

		if err != nil {
			return nil, err
		}

		post, err := service.GetPostByID(insertId)

		if err != nil {
			return nil, err
		}

		return post, nil
	},
}

var addCommentMutation = &graphql.Field{
	Type: CommentType,
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

		name := params.Args["name"].(string)
		content := params.Args["content"].(string)
		postID := params.Args["postId"].(int)

		insertID, err := service.AddComment(name, content, postID)

		if err != nil {
			return nil, err
		}

		comment, err := service.GetCommentByID(insertID)

		if err != nil {
			return nil, err
		}

		return comment, nil
	},
}
