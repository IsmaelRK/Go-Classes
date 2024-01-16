package main

import "fmt"

func main() {

	user := map[string]string{ // Key type and then value type
		"name":     "Ismael",
		"lastName": "R",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	mapWithMap := map[string]map[string]string{
		"name": {
			"name":   "Ismael",
			"lastNm": "R",
		},

		"name2": {
			"name":   "Ismael2",
			"lastNm": "R2",
		},

		// more key/values,
	}

	fmt.Println(mapWithMap)
	fmt.Println(mapWithMap["name"])
	fmt.Println(mapWithMap["name"]["name"])
	fmt.Println(mapWithMap["name"]["lastNm"])

	delete(mapWithMap, "name2")
	fmt.Println(mapWithMap)

	for key, submap := range mapWithMap {
		fmt.Println(key)
		fmt.Println(submap)

		delete(submap, "lastNm")
	}

	fmt.Println(mapWithMap)
}
