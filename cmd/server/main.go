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

	_ "github.com/lib/pq"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	env, err := app.ENV()
	if err != nil {
		return err
	}

	connectionString := env["PSQL_URI"]

	db, err := setupDatabase(connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

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
	userService := api.NewUserService(storage)

	server := app.NewServer(router, postService, userService)

	server.Run()

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

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS posts (id SERIAL, title VARCHAR(255) NOT NULL UNIQUE, author_id INT NOT NULL, md_body TEXT NOT NULL, created_at DATE NOT NULL, updated_at DATE NOT NULL)")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL, email VARCHAR(255) NOT NULL UNIQUE, password VARCHAR(255) NOT NULL, role_id int NOT NULL, username VARCHAR(255) NOT NULL, first_name VARCHAR(255) NOT NULL, last_name VARCHAR(255) NOT NULL, created_at DATE NOT NULL, updated_at DATE NOT NULL)")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS roles (id SERIAL, role_name VARCHAR(255) NOT NULL UNIQUE)")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS comments (id SERIAL, comment_body TEXT NOT NULL, user_id INT, post_id INT, created_at DATE NOT NULL, updated_at DATE NOT NULL)")
	if err != nil {
		return nil, err
	}

	return db, nil
}
