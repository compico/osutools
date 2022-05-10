package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var html string = "./static/html/"

func appHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles(html+"head.html", html+"footer.html", html+"app.html")
	if err != nil {
		fmt.Fprintf(w, "Error \"ParseFiles\": %v\n", err.Error())
	}
	if err := t.ExecuteTemplate(w, "app", nil); err != nil {
		fmt.Fprintf(w, "Error \"ExecuteTemplate\": %v\n", err.Error())
	}
}
