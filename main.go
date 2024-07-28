package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)

	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	return acc
}

func seedAccounts(store Storage) {
	seedAccount(store, "Zedan", "Mohamed", "HunterxHunter")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")

	flag.Parse()

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
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	// seed stuff
	if *seed {
		fmt.Println("Seeding the database")
		seedAccounts(store)
	}

	apiServer := NewAPIServer(listenAddr, store)

	apiServer.Serve()

}
