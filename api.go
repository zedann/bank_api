package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIServer struct {
	ListenAddr string
	Store      Storage
}

var apiServer *APIServer

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	apiServer = &APIServer{
		ListenAddr: listenAddr,
		Store:      store,
	}
	return apiServer
}

func GetTheApiServer() *APIServer {
	return apiServer
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()

	router.HandleFunc("/login", HTTPHandleFunc(HandleLogin)).Methods("POST")
	router.HandleFunc("/accounts", HTTPHandleFunc(HandleGetAccounts)).Methods("GET")
	router.HandleFunc("/accounts/{id}", WithJWTAuth(HTTPHandleFunc(HandleGetAccountByID))).Methods("GET")
	router.HandleFunc("/accounts", HTTPHandleFunc(HandleCreateAccount)).Methods("POST")
	router.HandleFunc("/accounts/{id}", HTTPHandleFunc(HandleDeleteAccount)).Methods("DELETE")
	router.HandleFunc("/transfer", HTTPHandleFunc(HandleTransfer)).Methods("POST")

	fmt.Println("JSON API RUN ON SERVER AT ", s.ListenAddr)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allowing all origins, you should change this to specific origins in production
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		Debug:          true, // Enable debug mode to see logs
	})

	// Wrap the router with the CORS handler
	handler := c.Handler(router)
	http.ListenAndServe(s.ListenAddr, handler)
}
