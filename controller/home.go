package controller

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("home.html").ParseFiles("./frontend/home.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}
