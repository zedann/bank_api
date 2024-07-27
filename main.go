package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	listenAddr := os.Getenv("LISTEN_ADDR")
	store, err := NewPostgresStore(PostgresConfig{
		DbUser:     os.Getenv("DB_USER"),
		DbName:     os.Getenv("DB_NAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatal(err)
	}

	apiServer := NewAPIServer(listenAddr, store)

	apiServer.Serve()

}
