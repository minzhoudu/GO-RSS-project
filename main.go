package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/minzhoudu/rss-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfiguration struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the ENV")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the ENV")
	}

	connection, sqlErr := sql.Open("postgres", dbUrl)
	if sqlErr != nil {
		log.Fatal("Unable to connect to the database")
	}

	apiConfig := apiConfiguration{
		DB: database.New(connection),
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	//Health
	v1Router.Get("/healthz", handleReadiness)
	v1Router.Get("/error", handleError)

	//Users
	v1Router.Post("/users", apiConfig.handlerCreateUser)
	v1Router.Get("/users", apiConfig.middlewareAuth(apiConfig.handleGetUserByApiKey))

	//Feed
	v1Router.Post("/feeds", apiConfig.middlewareAuth(apiConfig.handleCreateFeed))
	v1Router.Get("/feeds", apiConfig.handleGetFeeds)

	//Feed Follows
	v1Router.Post("/feed-follow", apiConfig.middlewareAuth(apiConfig.handlerCreateFeedFollow))
	v1Router.Get("/feed-follow", apiConfig.middlewareAuth(apiConfig.handlerGetFeedFollows))
	v1Router.Delete("/feed-follow/{feedFollowId}", apiConfig.middlewareAuth(apiConfig.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port: %v", portString)
	srvErr := server.ListenAndServe()
	if srvErr != nil {
		log.Fatal(srvErr)
	}
}
