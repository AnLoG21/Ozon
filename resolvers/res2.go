package resolvers

import (
	"awesomeProject/database"
	"awesomeProject/models"
	_ "context"
	_ "fmt"
	"github.com/graphql-go/graphql"
)

var postType = graphql.NewObject(
	graphql.ObjectConfig{
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
			"allowComments": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var (
	PostQuery = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "PostQuery",
			Fields: graphql.Fields{
				"post": &graphql.Field{
					Type:        postType,
					Description: "Get post by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						id, _ := params.Args["id"].(int)
						post, err := database.GetPostByID(id)
						return post, err
					},
				},
			},
		},
	)
)

var CreatePostMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CreatePostMutation",
		Fields: graphql.Fields{
			"createPost": &graphql.Field{
				Type:        postType,
				Description: "Create a new post",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"allowComments": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					post := models.Post{
						Title:         params.Args["title"].(string),
						Content:       params.Args["content"].(string),
						AllowComments: params.Args["allowComments"].(bool),
					}
					err := database.CreatePost(post)
					return post, err
				},
			},
		},
	},
)
