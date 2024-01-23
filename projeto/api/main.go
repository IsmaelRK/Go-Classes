package main

import (
	"api/src/config"
	"api/src/router"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

func init() {

	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	string64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(string64)

}

func main() {

	config.LoadEnv()
	fmt.Println(config.Port)

	fmt.Println("Api is UP")

	r := router.GenRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
