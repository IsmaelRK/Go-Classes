package main

import (
	"fmt"
	"testing"
)

type test struct {
	insert      string
	returnValue string
}

func TestAddrsType(t *testing.T) {

	mytests := []test{
		{"Rua 1", "rua"},
		{"Avenida 1", "avenida"},
		{"Logradouro 378", "not found"},
	}

	for _, key := range mytests {

		fmt.Println(key)

		addressTypeReceived := Addrstype(key.insert)

		fmt.Println(key.insert, addressTypeReceived)

		if addressTypeReceived != key.returnValue {
			t.Errorf("The type is different!")
		}

	}

}

//func TestAddrsType(t *testing.T) {
//
//	tstAddres := "Avenida XYZ"
//	addrWaited := "avenida"
//
//	addressTypeReceived := Addrstype(tstAddres)
//
//	fmt.Println(addrWaited, addressTypeReceived)
//
//	if addrWaited != addressTypeReceived {
//		t.Errorf("The type is different!")
//	}
//
//}
