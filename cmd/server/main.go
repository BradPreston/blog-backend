package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/BradPreston/blog-backend/pkg/api"
	"github.com/BradPreston/blog-backend/pkg/app"
	"github.com/BradPreston/blog-backend/pkg/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	env, err := godotenv.Read(".env")
	if err != nil {
		return err
	}

	connectionString := env["PSQL_URI"]

	db, err := setupDatabase(connectionString)
	if err != nil {
		return err
	}

	storage := repository.NewStorage(db)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	postService := api.NewPostService(storage)

	server := app.NewServer(router, postService)

	err = server.Run()
	if err != nil {
		return err
	}

	return nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
