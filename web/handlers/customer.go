package handlers

import (
	"Gapp/web/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type CustomerAPI struct {
}

type AuthResp struct {
	Token string `json:"jwtToken"`
	Algo  string `json:"algo"`
}

func (capi *CustomerAPI) Authenticate(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	json.NewDecoder(r.Body).Decode(&customer)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2023, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("venugopal"))

	authResp := &AuthResp{Token: tokenString, Algo: "SHA256"}
	fmt.Println(tokenString, err)
	json.NewEncoder(w).Encode(authResp)

	//fmt.Fprintf(w, "%s %s  years old!", customer.Firstname, customer.Lastname)
}

func (capi *CustomerAPI) SignUp(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	fmt.Fprintf(w, "%s %s  years old!", customer.Firstname, customer.Lastname)
}

func (capi *CustomerAPI) CustInfo(w http.ResponseWriter, r *http.Request) {
	contact := models.ContactInfo{Age: 25, Email: "user@ecommerce.com"}
	customer := models.Customer{
		Firstname:   "John",
		Lastname:    "Doe",
		ContactInfo: contact,
	}
	json.NewEncoder(w).Encode(customer)
	//fmt.Fprintf(w, "Hello, world = %s", "Coder")
}
