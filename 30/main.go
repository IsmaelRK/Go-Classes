package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {

	//w.Write([]byte("<h1>Hello World</h1>"))

	u := user{
		"Ismael",
		"ab@gmail.com",
	}

	templates.ExecuteTemplate(w, "index.html", u)

}

func main() {

	templates = template.Must(template.ParseGlob("*html"))

	http.HandleFunc("/home", home)

	fmt.Println("Listening on :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
