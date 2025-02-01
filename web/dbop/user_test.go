package dbop

import (
	"Gapp/web/models"
	"fmt"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	IntializePool()
	defer ClosePool()
	euser, authError := Authenticate(models.AuthRequest{UserName: "venugopal@ecom.com", Password: "ecom#24"})
	if authError == nil {
		fmt.Println(euser)
	} else {
		t.Fail()
	}

}
