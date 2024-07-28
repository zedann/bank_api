package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func CreateJWT(account *Account) (string, error) {

	// Create the Claims
	claims := &jwt.MapClaims{
		// expires in 30 days
		"ExpiresAt":     jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		"AccountNumber": account.Number,
	}

	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func PermissionDenied(w http.ResponseWriter) {
	WriteJson(w, http.StatusForbidden, APIError{Error: "permission denied"})
}

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling jwt auth middleware")

		tokenStr := r.Header.Get("x-jwt-token")
		token, err := ValidateJWT(tokenStr)

		if err != nil {
			PermissionDenied(w)
			return
		}
		if !token.Valid {
			PermissionDenied(w)
			return
		}
		userID, err := getID(r)
		if err != nil {
			PermissionDenied(w)
			return
		}

		storage := GetTheApiServer().Store
		account, err := storage.GetAccountByID(userID)

		if err != nil {
			PermissionDenied(w)
			return
		}

		// casting
		claims := token.Claims.(jwt.MapClaims)

		// fmt.Printf("claims acc type : %T , accNum : %T", claims["AccountNumber"], account.Number)
		if account.Number != int64(claims["AccountNumber"].(float64)) {
			PermissionDenied(w)
			return
		}

		handlerFunc(w, r)
	}

}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
}

func HTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, APIError{
				Error: err.Error(),
			})
		}

	}
}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil

}
