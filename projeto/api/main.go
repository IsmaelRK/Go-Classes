package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Api is UP")

	r := router.GenRouter()

	log.Fatal(http.ListenAndServe(":5000", r))
}
