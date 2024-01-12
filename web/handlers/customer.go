package handlers

import (
	"Gapp/web/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type CustomerAPI struct {
}

func (capi *CustomerAPI) SignUp(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	fmt.Fprintf(w, "%s %s  years old!", customer.Firstname, customer.Lastname)
}

func (capi *CustomerAPI) SignIn(w http.ResponseWriter, r *http.Request) {
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
