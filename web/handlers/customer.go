package handlers

import (
	"Gapp/web/dbop"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerAPI struct {
}

func (capi *CustomerAPI) CustInfo(c *gin.Context) {
	customer := dbop.GetUserInfo(1)
	//json.NewEncoder(w).Encode(customer)
	c.JSON(http.StatusOK, customer)
	//fmt.Fprintf(w, "Hello, world = %s", "Coder")
}
