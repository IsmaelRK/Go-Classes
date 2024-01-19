package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	fmt.Println("Listening on 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
