package resolvers

import (
	"awesomeProject/database"
	"awesomeProject/models"
	_ "context"
	"github.com/graphql-go/graphql"
)

var commentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"postId": &graphql.Field{
				Type: graphql.Int,
			},
			"parentId": &graphql.Field{
				Type: graphql.Int,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"authorId": &graphql.Field{
				Type: graphql.Int,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var commentQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CommentQuery",
		Fields: graphql.Fields{
			"comments": &graphql.Field{
				Type:        graphql.NewList(commentType),
				Description: "Get comments for a post",
				Args: graphql.FieldConfigArgument{
					"postId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					postId, _ := params.Args["postId"].(int)
					comments, err := database.GetCommentsByPostID(postId)
					return comments, err
				},
			},
		},
	},
)

var createCommentMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CreateCommentMutation",
		Fields: graphql.Fields{
			"createComment": &graphql.Field{
				Type:        commentType,
				Description: "Create a new comment",
				Args: graphql.FieldConfigArgument{
					"postId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"parentId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"authorId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					comment := models.Comment{
						PostID:   params.Args["postId"].(int),
						ParentID: params.Args["parentId"].(*int),
						Content:  params.Args["content"].(string),
						AuthorID: params.Args["authorId"].(int),
					}
					err := database.CreateComment(comment)
					return comment, err
				},
			},
		},
	},
)
