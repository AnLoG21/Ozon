package subscriptions

import (
	"github.com/graphql-go/graphql"
)

var (
	commentType = graphql.NewObject(graphql.ObjectConfig{})
)

var CommentAddedSubscriptionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CommentAddedSubscription",
	Fields: graphql.Fields{
		"comment": &graphql.Field{
			Type: CommentType, // Предполагается, что у вас есть тип CommentType
		},
	},
})

// CommentAddedSubscription - функция-разрешитель для подписки на добавление комментариев
var CommentAddedSubscription = func(p graphql.ResolveParams) (interface{}, error) {
	// Логика для получения добавленного комментария
	// Это может быть через каналы, WebSockets, PubSub систему или любой другой механизм
	// Для примера, возвращаем фиктивный комментарий
	return map[string]interface{}{
		"comment": map[string]interface{}{
			"id":      1,
			"post_id": 1,
			"content": "Новый комментарий",
			// ... другие поля комментария
		},
	}, nil
}

// Subscription - корневой тип подписок
var Subscription = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		"commentAdded": &graphql.Field{
			Type:    CommentAddedSubscriptionType,
			Resolve: CommentAddedSubscription,
			Args: graphql.FieldConfigArgument{
				"postID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
		},
	},
})

// Schema - корень GraphQL схемы
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Subscription: Subscription,
})

// CommentType - тип GraphQL для модели Comment
var CommentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"post_id": &graphql.Field{
			Type: graphql.Int,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		// ... другие поля комментария
	},
})
