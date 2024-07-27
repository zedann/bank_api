package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

	router.HandleFunc("/accounts", HTTPHandleFunc(HandleGetAccounts)).Methods("GET")
	router.HandleFunc("/accounts/{id}", HTTPHandleFunc(HandleGetAccountByID)).Methods("GET")
	router.HandleFunc("/accounts", HTTPHandleFunc(HandleCreateAccount)).Methods("POST")
	router.HandleFunc("/accounts/{id}", HTTPHandleFunc(HandleDeleteAccount)).Methods("DELETE")

	fmt.Println("JSON API RUN ON SERVER AT ", s.ListenAddr)

	http.ListenAndServe(s.ListenAddr, router)
}
