package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {

	//w.Write([]byte("<h1>Hello World</h1>"))
	templates.ExecuteTemplate(w, "index.html", nil)

}

func main() {

	templates = template.Must(template.ParseGlob("*html"))

	http.HandleFunc("/home", home)

	fmt.Println("Listening on :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
