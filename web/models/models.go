package models

type ContactInfo struct {
	Age   int    `json:"age"`
	Email string `json:"eamil"`
}

type Customer struct {
	Firstname   string      `json:"firstname"`
	Lastname    string      `json:"lastname"`
	ContactInfo ContactInfo `json:"contact"`
}

type Seller struct {
	Firstname   string      `json:"firstname"`
	Lastname    string      `json:"lastname"`
	ContactInfo ContactInfo `json:"contact"`
}

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}
