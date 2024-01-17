package main

import (
	"cli/app"
	"log"
	"os"
)

func main() {

	application := app.Generate()
	erro := application.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}

}