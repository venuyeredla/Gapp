package handlers

import (
	"Gapp/web/dbop"
	"Gapp/web/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerAPI struct {
}

func (capi *CustomerAPI) CustInfo(c *gin.Context) {
	customer, authError := dbop.Authenticate(models.AuthRequest{})
	if authError != nil {
		// json.NewEncoder(w).Encode(customer)
		c.JSON(http.StatusOK, customer)
		//fmt.Fprintf(w, "Hello, world = %s", "Coder")
	}

}
