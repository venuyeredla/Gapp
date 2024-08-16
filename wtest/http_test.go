package wtest

import (
	"Gapp/web/handlers"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
)

const (
	DOMAIN_NAME = "http://localhost:2024"
)

func TestAuthEndpoint(t *testing.T) {
	/*
			{
		    "username":"",
		    "password" : ""
		    }
	*/

	requesBody := &handlers.AuthRequest{UserName: "venugopal@ecom.com", Password: "ecom#24"}
	bytearr, error := json.Marshal(requesBody)

	if error != nil {
		log.Default().Println("Error in marshalling strcut", error.Error())
	}
	req, _ := http.NewRequest("POST", DOMAIN_NAME+"/api/auth", bytes.NewBuffer(bytearr))
	response, h_error := http.DefaultClient.Do(req)
	if h_error == nil {
		bodyBytes, _ := io.ReadAll(response.Body)
		var resp handlers.AuthResponse
		json.Unmarshal(bodyBytes, &resp)
		log.Default().Println(resp)

	} else {
		log.Default().Println("Error in connecting api")
	}
	defer response.Body.Close()

}
