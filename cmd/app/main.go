package main

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/zedann/bank_api/api"
)
func main() {

	godotenv.Load()

	listenAddr := os.Getenv("LISTEN_ADDR")

	apiServer := api.NewServer(listenAddr)

	apiServer.Run()
}
