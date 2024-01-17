package main

import (
	"fmt"
	"strings"
)

func addrstype(adress string) string {

	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	adress = strings.ToLower(adress)
	firtWord := strings.Split(adress, " ")[0]

	var resultType string

	for _, validType := range validTypes {

		if firtWord == validType {
			resultType = validType
			break
		} else {
			resultType = "Not Found"
		}

	}

	return strings.ToTitle(resultType)
}

func main() {
	typeAdrrs := addrstype("rua 1")
	fmt.Println(typeAdrrs)
}
