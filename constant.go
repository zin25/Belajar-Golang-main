package main

// Constanta Nilai tidak berubah

import "fmt"

func main() {
	const firstName string = "Azzin"
	const lastname string = "Maharil"

	fmt.Println(firstName+"", lastname)

	// multi constant

	const (
		umur           = 23
		panjang string = "Panjangnya 22cm"
	)

	fmt.Println(umur)
	fmt.Println(panjang)

}
