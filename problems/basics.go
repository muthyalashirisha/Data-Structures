package problems

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDB() {

	mongouri := os.Getenv("MONGOURI")

	//mongo connection
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongouri))
	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
}
func api() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error while loading env file", err)
	}
	args := os.Args
	port := os.Getenv("PORT")
	if len(args) > 0 {
		port = args[1]
	}
	server := &http.Server{
		Addr:         ":" + port,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
		Handler:      httpServer(),
	}
	fmt.Println(server)
}
func httpServer() http.Handler {
	var handler http.Handler
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/v1").Subrouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
	handler = c.Handler(subRouter)
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.Header().Add("Accept-Encoding", "gzip")
			next.ServeHTTP(w, r)
		})
	})
	subRouter.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("<html>")
	})
	http.Handle("/", router)
	return handler
}
