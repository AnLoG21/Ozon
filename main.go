package main

import (
	"awesomeProject/resolvers"
	"awesomeProject/subscriptions"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

func main() {
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    resolvers.PostQuery,
			Mutation: resolvers.CreatePostMutation,
			Subscription: graphql.SubscriptionConfig{
				Resolve: subscriptions.CommentAddedSubscription,
			},
		},
	)

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
