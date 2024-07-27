package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

func HandleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	return errors.New("NOT FOUND")
}
func HandleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	apiServer := GetTheApiServer()

	account, err := apiServer.Store.GetAccountByID(id)

	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, account)

}
func HandleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	apiServer := GetTheApiServer()

	createdAccReq := new(CreateAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(createdAccReq); err != nil {
		return err
	}

	account, err := NewAccount(createdAccReq.FirstName, createdAccReq.LastName)

	if err != nil {
		return err
	}

	if err := apiServer.Store.CreateAccount(account); err != nil {
		return err
	}

	return WriteJson(w, http.StatusCreated, account)
}
func HandleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func HandleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
