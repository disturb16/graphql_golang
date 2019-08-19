package schema

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/handler"

	"github.com/disturb16/graphql_golang/schema/mutations"
	"github.com/disturb16/graphql_golang/schema/queries"
	"github.com/disturb16/graphql_golang/settings"
)

// NewHandler main handler for graphql
func NewHandler() (*handler.Handler, error) {
	config, err := settings.Configuration("./")
	if err != nil {
		log.Fatal(err)
	}

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"post":   queries.PostQuery,
			"posts":  queries.PostsQuery,
			"author": queries.AuhtorQuery,
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "mutation",
		Fields: graphql.Fields{
			"add_post":    mutations.AddPostMutation,
			"add_comment": mutations.AddCommentMutation,
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	if err != nil {
		return nil, err
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		GraphiQL: config.UseGraphiQL,
		FormatErrorFn: func(err error) gqlerrors.FormattedError {
			var formatted gqlerrors.FormattedError
			switch err := err.(type) {
			case *gqlerrors.Error:
				formatted = gqlerrors.FormatError(err)
			default:
				log.Println(err)
			}

			log.Println(err)
			return formatted
		},
	})

	return h, nil
}
