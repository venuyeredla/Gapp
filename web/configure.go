package web

import (
	"Gapp/web/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func ConfigureHttp() {
	log.Println("Setting up file server and api handlers")
	//Static file handler
	fileHandler := http.FileServer(http.Dir("webapp/"))
	http.Handle("/webapp/", http.StripPrefix("/webapp/", fileHandler))
	//opening introduction page.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "webapp/index.html")
	})

	//Generating entitre HTML using libraries
	sc := &handlers.StatiContent{}
	sc.Preprocess()
	http.HandleFunc("/form", sc.GenrateForm)

	var customerAPI handlers.CustomerAPI
	http.HandleFunc("/signup", logging(customerAPI.SignUp))
	http.HandleFunc("/custinfo", logging(customerAPI.CustInfo))

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

}
