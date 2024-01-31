package handlers

import (
	"Gapp/web/models"
	"html/template"
	"net/http"
)

type StatiContent struct {
	TemplateMap map[string]*template.Template
}

func (sc *StatiContent) Preprocess() {
	sc.TemplateMap = make(map[string]*template.Template)
	//sc.TemplateMap["todo"] = template.Must(template.ParseFiles("webapp/html/layout.html"))
	sc.TemplateMap["formtmpl"] = template.Must(template.ParseFiles("wstatic/form.html"))
}

func (sc *StatiContent) GenrateForm(w http.ResponseWriter, r *http.Request) {
	/*if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	} */

	details := models.ContactInfo{
		Email: r.FormValue("email"),
	}
	formtmpl := sc.TemplateMap["formtmpl"]

	// do something with details
	_ = details

	data := models.TodoPageData{
		PageTitle: "My TODO list",
		Todos: []models.Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	d := struct {
		Success bool
		Todo    models.TodoPageData
	}{Success: false, Todo: data}

	formtmpl.Execute(w, d)
}
