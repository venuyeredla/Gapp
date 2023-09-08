// App project main.go
package main

import (
	"Gapp/webapi"
	"fmt"
	"net/http"
	"os"
)

func main() {
	host, _ := os.Hostname()
	fmt.Printf("Host %s, Page/Block size =%v \n", host, os.Getpagesize())

	webapi.ConfigureHttp()
	fmt.Println("Http Server running on port. 2023")
	http.ListenAndServe(":2023", nil)
}
