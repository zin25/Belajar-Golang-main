package main

import "fmt"

func ChangeAddressToIndonesia(address Address) {
	address.Country = "Indonesia"
}

func main() {
	var address = Address{"Lampung", "Bandar Lampung", ""}
	ChangeAddressToIndonesia(address)

	fmt.Println(address)
}
