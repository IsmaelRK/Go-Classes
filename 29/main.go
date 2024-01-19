package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
}

func main() {

	dogJSON := `{"nome":"bigdog","raca":"yorkshire"}`

	var d dog

	err := json.Unmarshal([]byte(dogJSON), &d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dogJSON)

	d2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(dogJSON), &d2) {
		
	}

}
