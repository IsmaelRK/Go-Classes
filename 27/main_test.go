package main

import (
	"fmt"
	"testing"
)

func TestAddrsType(t *testing.T) {

	tstAddres := "Avenida XYZ"
	addrWaited := "avenida"

	addressTypeReceived := Addrstype(tstAddres)

	fmt.Println(addrWaited, addressTypeReceived)

	if addrWaited != addressTypeReceived {
		t.Errorf("The type is different!")
	}

}
