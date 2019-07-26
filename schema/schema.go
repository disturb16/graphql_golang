package schema

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/handler"

	"github.com/disturb16/graphql_golang/internal/services"
	"github.com/disturb16/graphql_golang/settings"
)

type graphqlSchema struct {
	Service *services.Service
}

var service *services.Service

// New main handler for graphql
func New(s *services.Service) (*handler.Handler, error) {
	config, err := settings.Configuration("./")
	if err != nil {
		log.Fatal(err)
	}

	service = s

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"post":   postQuery,
			"posts":  postsQuery,
			"author": auhtorQuery,
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "mutation",
		Fields: graphql.Fields{
			"add_post":    addPostMutation,
			"add_comment": addCommentMutation,
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
