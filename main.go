package main

import (
	"awesomeProject/resolvers"
	_ "awesomeProject/subscriptions"
	"context"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    resolvers.PostQuery,
			Mutation: resolvers.CreatePostMutation,
		},
	)

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("localhost/graphql", h)
	http.ListenAndServe("http://localhost", nil)
	srv := &http.Server{
		Addr:    ":5432",
		Handler: http.DefaultServeMux,
	}

	go func() {
		log.Println("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
