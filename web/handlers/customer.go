package handlers

import (
	"Gapp/web/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

type CustomerAPI struct {
}

type AuthResp struct {
	Token string `json:"jwtToken"`
	Algo  string `json:"algo"`
}

func (capi *CustomerAPI) Authenticate(c *gin.Context) {
	// var customer models.Customer
	//json.NewDecoder(r.Body).Decode(&customer)

	//val, ok := c.Params.Get("name")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2023, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("venugopal"))

	authResp := &AuthResp{Token: tokenString, Algo: "SHA256"}
	fmt.Println(tokenString, err)
	//json.NewEncoder(w).Encode(authResp)

	c.JSON(http.StatusOK, authResp)

	//fmt.Fprintf(w, "%s %s  years old!", customer.Firstname, customer.Lastname)
}

func (capi *CustomerAPI) SignUp(c *gin.Context) {
	var customer models.Customer
	json.NewDecoder(c.Request.Body).Decode(&customer)
	c.String(http.StatusOK, "Success fully signed up", customer.Firstname, customer.Lastname)
}

func (capi *CustomerAPI) CustInfo(c *gin.Context) {
	contact := models.ContactInfo{Age: 25, Email: "user@ecommerce.com"}
	customer := models.Customer{
		Firstname:   "John",
		Lastname:    "Doe",
		ContactInfo: contact,
	}
	//json.NewEncoder(w).Encode(customer)
	c.JSON(http.StatusOK, customer)
	//fmt.Fprintf(w, "Hello, world = %s", "Coder")
}
