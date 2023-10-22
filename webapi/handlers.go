package webapi

import (
	"Gapp/types"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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

func ConfigureHttp() {
	fs := http.FileServer(http.Dir("webapp/"))

	http.Handle("/webapp/", http.StripPrefix("/webapp/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "webapp/index.html")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world = %s", "Coder")
	})

	tmpl := template.Must(template.ParseFiles("webapp/html/layout.html"))
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

	formtmpl := template.Must(template.ParseFiles("webapp/html/form.html"))
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

	http.HandleFunc("/log", logging(foo))
	http.HandleFunc("/encode", logging(UserInfo))
	http.HandleFunc("/decode", logging(PostUserInfo))
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	user1 := types.User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
	}
	json.NewEncoder(w).Encode(user1)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}
