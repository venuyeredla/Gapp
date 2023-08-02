// App project main.go
package main

import (
	"Gapp/types"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	host, _ := os.Hostname()
	fmt.Printf("Host %s, Page/Block size =%v \n", host, os.Getpagesize())

	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	tmpl := template.Must(template.ParseFiles("static/layout.html"))
	http.HandleFunc("/tmpl", func(w http.ResponseWriter, r *http.Request) {
		data := types.TodoPageData{
			PageTitle: "My TODO list",
			Todos: []types.Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	formtmpl := template.Must(template.ParseFiles("static/form.html"))
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		/*if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		} */

		details := types.ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		_ = details

		formtmpl.Execute(w, struct{ Success bool }{false})
	})

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

	http.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/websocket.html")
	})

	http.HandleFunc("/log", logging(foo))
	http.HandleFunc("/encode", logging(UserInfo))
	http.HandleFunc("/decode", logging(PostUserInfo))

	fmt.Println("Http Server running on port. 2023")
	http.ListenAndServe(":2023", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	user1 := types.User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
	}
	json.NewEncoder(w).Encode(user1)
}

func PostUserInfo(w http.ResponseWriter, r *http.Request) {
	var user types.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}
