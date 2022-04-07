package main

import (
	"fmt"
)

func variable() {
	var nama string

	nama = "Azzin Maharil"
	fmt.Println(nama)

	nama = "Azzin Maharil Ganteng"
	fmt.Println(nama)

	var friendName = "budi" // Tanpa tipe data cth String
	fmt.Println(friendName)

	var age = 22
	fmt.Println(age)

	negara := "Indonesia" // tanpa var, biar lebih simple
	fmt.Println(negara)

	var (
		firstname = "Azzin" //Multi Variable
		lastname  = "Maharil"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)

}
