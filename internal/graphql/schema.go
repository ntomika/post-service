package graphql

import (
	"github.com/graphql-go/graphql"
	"post-service/internal/handlers"
)

func NewSchema(server *handlers.Server) (graphql.Schema, error) {
	// Define Comment type
	commentType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.ID},
			"post_id":    &graphql.Field{Type: graphql.ID},
			"parent_id":  &graphql.Field{Type: graphql.ID},
			"content":    &graphql.Field{Type: graphql.String},
			"author_id":  &graphql.Field{Type: graphql.ID},
			"created_at": &graphql.Field{Type: graphql.String},
		},
	})

	// Define Post type
	postType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id":             &graphql.Field{Type: graphql.ID},
			"title":          &graphql.Field{Type: graphql.String},
			"content":        &graphql.Field{Type: graphql.String},
			"author_id":      &graphql.Field{Type: graphql.ID},
			"allow_comments": &graphql.Field{Type: graphql.Boolean},
			"created_at":     &graphql.Field{Type: graphql.String},
			"comments":       &graphql.Field{Type: graphql.NewList(commentType)},
		},
	})

	// Define Query type
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"posts": &graphql.Field{
				Type:    graphql.NewList(postType),
				Resolve: server.GetPostList,
			},
			"post": &graphql.Field{
				Type: postType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				Resolve: server.GetPost,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
				Args: graphql.FieldConfigArgument{
					"post_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
					"limit":   &graphql.ArgumentConfig{Type: graphql.Int},
					"offset":  &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: server.GetCommentList,
			},
			"comment": &graphql.Field{
				Type: commentType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				Resolve: server.GetComment,
			},
		},
	})

	// Define Mutation type
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createPost": &graphql.Field{
				Type: postType,
				Args: graphql.FieldConfigArgument{
					"title":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"content":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"author_id":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
					"allow_comments": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Boolean)},
				},
				Resolve: server.CreatePost,
			},
			"createComment": &graphql.Field{
				Type: commentType,
				Args: graphql.FieldConfigArgument{
					"post_id":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
					"parent_id": &graphql.ArgumentConfig{Type: graphql.ID},
					"content":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"author_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				Resolve: server.CreateComment,
			},
		},
	})

	// Define schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		return graphql.Schema{}, err
	}

	return schema, nil
}
