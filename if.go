package main

import "fmt"

func main() {

	name := "maharil"

	if name == "azzin" {
		fmt.Println("Hallo Azzin")
	} else if name == "Joko" {
		fmt.Println("Hai Joko")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	if lenght := len(name); lenght > 4 { // Short Statement
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah benar")
	}

}
