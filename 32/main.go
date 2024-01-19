package main

import (
	"crud/server"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/user", server.CreateUsr).Methods(http.MethodPost)

	fmt.Println("Listening on 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
