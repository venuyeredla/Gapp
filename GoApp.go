// App project main.go
package main

import (
	"Gapp/web"
	"log"
	"net/http"
	"os"
)

const (
	port = 2023
)

func main() {
	host, _ := os.Hostname()
	// fmt.Printf("Host %s, Page/Block size =%v \n", host, os.Getpagesize())
	web.ConfigureHttp()
	log.Printf("Application is available at Host: %v and  Port=%v", host, port)
	//log.P("Http Server running on port. 2023")
	http.ListenAndServe(":2023", nil)

}
