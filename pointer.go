package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) { // jadi pointer *
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Bandar Lampung", "Lampung", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Jakarta"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var alamat = Address{"Bandar Lampung", "Lampung", ""}
	ChangeAddressToIndonesia(&alamat) // panggil pointer dengan &

	fmt.Println(alamat)
}
