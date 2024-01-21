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
	router.HandleFunc("/user", server.SearchUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.SearchUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.UpdateUser).Methods(http.MethodPut)

	fmt.Println("Listening on localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
