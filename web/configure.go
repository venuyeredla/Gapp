package web

import (
	"Gapp/web/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logging(ginHandler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Before = %v", c.Request.URL.Path)
		ginHandler(c)
		log.Printf("After = %v", c.Request.URL.Path)
	}
}

func JwtValidate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
		log.Println(r.URL.Path)
	}
}

func ApiConfig() *gin.Engine {
	log.Println("Setting up file server and api handlers")
	router := gin.Default()

	//Static file handler
	router.StaticFS("/wstatic/", http.Dir("wstatic/"))
	router.StaticFile("/", "wstatic/index.html")
	/*router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))
	*/
	var customerAPI handlers.CustomerAPI
	router.GET("/api/authenticate", Logging(customerAPI.Authenticate))
	router.POST("/api/signup", Logging(customerAPI.SignUp))
	router.GET("/api/custinfo/:id", Logging(customerAPI.CustInfo))

	//Generating entitre HTML using libraries
	sc := &handlers.StatiContent{}
	//sc.Preprocess()
	router.LoadHTMLGlob("wstatic/*")
	router.GET("/tmpl", Logging(sc.GenrateForm))

	return router
}

/*


func ConfigureHttp() {

	func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
		log.Println(r.URL.Path)
	}
}
	log.Println("Setting up file server and api handlers")
	//Static file handler
	fileHandler := http.FileServer(http.Dir("wstatic/"))
	http.Handle("/wstatic/", http.StripPrefix("/wstatic/", fileHandler))

	//opening introduction page.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "wstatic/index.html")
	})

	//Generating entitre HTML using libraries
	sc := &handlers.StatiContent{}
	sc.Preprocess()
	http.HandleFunc("/form", sc.GenrateForm)

	var customerAPI handlers.CustomerAPI

	http.HandleFunc("/api/authenticate", logging(customerAPI.Authenticate))
	http.HandleFunc("/signup", logging(customerAPI.SignUp))
	http.HandleFunc("/custinfo", logging(customerAPI.CustInfo))


	http.ListenAndServe(":2023", nil)

}



var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


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

*/
