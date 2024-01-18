package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
}

func main() {

	d := dog{"bigdog", "yorkshire"}
	fmt.Println(d)

	dogToJSON, erro := json.Marshal(d)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(dogToJSON)
	fmt.Println(bytes.NewBuffer(dogToJSON))

	d2 := map[string]string{
		"nome": "theDogo",
		"raca": "indefinida",
	}

	dog2ToJSON, erro := json.Marshal(d2)
	fmt.Println(bytes.NewBuffer(dog2ToJSON))

}
