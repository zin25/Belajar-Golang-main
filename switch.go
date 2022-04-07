package main

import "fmt"

func main() {

	name := "Azin"

	switch name {
	case "Azin":
		fmt.Println("Hallo Azin")
	case "Laras":
		fmt.Println("Hallo Baras")
	default:
		fmt.Println("Kenalan dong")
	}

	// switch length := len(name); length < 5 {
	// case true:
	// 	fmt.Println("Nama Sudah benar")
	// case false:
	// 	fmt.Println("Nama terlalu panjang")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
