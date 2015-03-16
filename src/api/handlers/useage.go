package handlers

import (
	html "html/template"
	"net/http"
)

type Page struct {
	Title string
	Body  string
}

var useageTemplate *html.Template

var err error

func init() {
	//Get the html template
	useageTemplate, err = html.ParseFiles("templates/useage.html")
	if err != nil {
		panic(err)
	}
}

//show useage page for http://localhost:8080
func Useage(w http.ResponseWriter, r *http.Request) {
	page := Page{Title: "Useage", Body: "Useage"}
	useageTemplate.Execute(w, page)
}
