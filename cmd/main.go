package main

import (
	"context"
	"github.com/joho/godotenv"
	"os"
	"post-service/internal/graphql"
	"post-service/internal/handlers"
	"post-service/internal/repository"

	"log"
	"net/http"
	// "github.com/gorilla/websocket"
	"github.com/graphql-go/handler"
)

func main() {
	ctx := context.Background()

	var repo repository.Repository
	var err error

	err = godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Failed get ENVS: %v", err)
	}

	storageType := os.Getenv("STORAGE_TYPE")
	switch storageType {
	case "postgres":
		dsn := os.Getenv("DB_URL")
		repo, err = repository.NewPostgresRepository(ctx, dsn)
	case "inmemory":
		repo = repository.NewInMemoryRepository()
	default:
		log.Fatalf("Unknown storage type: %s", storageType)
	}

	server := handlers.NewServer(repo)
	schema, err := graphql.NewSchema(server)
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
