package main

import (
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Number   int64  `json:"number"`
	Password string `json:"password"`
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

type APIError struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Number            int64     `json:"number"`
	EncryptedPassword string    `json:"-"`
	Balance           float64   `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName, password string) (*Account, error) {

	encpwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encpwd),
		Number:            int64(rand.Intn(1000000)),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
