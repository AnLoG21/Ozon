package main

import (
	"github.com/graphql-go/graphql"
)

var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.ID},
			"title": &graphql.Field{Type: graphql.String},
			"body":  &graphql.Field{Type: graphql.String},
		},
	},
)

var CommentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id":     &graphql.Field{Type: graphql.ID},
			"postId": &graphql.Field{Type: graphql.ID},
			"body":   &graphql.Field{Type: graphql.String},
		},
	},
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"post": &graphql.Field{
				Type: PostType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.ID},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetPostByID(p.Args["id"].(string))
				},
			},
			"comment": &graphql.Field{
				Type: CommentType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.ID},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetCommentByID(p.Args["id"].(string))
				},
			},
		},
	},
)

var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createPost": &graphql.Field{
				Type: PostType,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"body":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return CreatePost(p.Args["title"].(string), p.Args["body"].(string))
				},
			},
			"createComment": &graphql.Field{
				Type: CommentType,
				Args: graphql.FieldConfigArgument{
					"postId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
					"body":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return CreateComment(p.Args["postId"].(string), p.Args["body"].(string))
				},
			},
		},
	},
)

var SubscriptionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "Subscription",
		Fields: graphql.Fields{},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:        QueryType,
		Mutation:     MutationType,
		Subscription: SubscriptionType,
	},
)

func main() {
}

func GetPostByID(id string) (interface{}, error) {
	return nil, nil
}

func GetCommentByID(id string) (interface{}, error) {
	return nil, nil
}

// CreatePost - функция для создания поста
func CreatePost(title, body string) (interface{}, error) {
	return nil, nil
}

func CreateComment(postID, body string) (interface{}, error) {
	return nil, nil
}
