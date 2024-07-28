package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	loginReq := new(LoginRequest)

	if err := json.NewDecoder(r.Body).Decode(loginReq); err != nil {
		return err
	}

	acc, err := GetTheApiServer().Store.GetAccountByNumber(int(loginReq.Number))
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", acc)
	if err := bcrypt.CompareHashAndPassword([]byte(acc.EncryptedPassword), []byte(loginReq.Password)); err != nil {
		return err
	}

	// create jwt token and return it

	tokenStr, err := CreateJWT(acc)

	// TODO: Create Cookie and Send it to the client

	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, map[string]any{
		"status": "success",
		"token":  tokenStr,
		"data":   acc,
	})
}

func HandleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	apiServer := GetTheApiServer()

	accounts, err := apiServer.Store.GetAccounts(10)

	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, accounts)
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

	account, err := NewAccount(createdAccReq.FirstName, createdAccReq.LastName, createdAccReq.Password)

	if err != nil {
		return err
	}

	if err := apiServer.Store.CreateAccount(account); err != nil {
		return err
	}

	token, err := CreateJWT(account)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusCreated, map[string]any{
		"status": "success",
		"token":  token,
		"data":   account,
	})

}

func HandleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)

	if err != nil {
		return err
	}

	apiServer := GetTheApiServer()
	if err := apiServer.Store.DeleteAccount(id); err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, map[string]int{"deleted": id})
}

func HandleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferReq := new(TransferRequest)

	if err := json.NewDecoder(r.Body).Decode(transferReq); err != nil {
		return err
	}

	defer r.Body.Close()

	return WriteJson(w, http.StatusOK, transferReq)
}
