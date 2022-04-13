package main

import "fmt"

func main() {

	person := map[string]int{
		"age":   23,
		"tahun": 1999,
	}

	person["bulan"] = 04

	fmt.Println(person)
	fmt.Println(person["age"])
	fmt.Println(person["tahun"]) // Ambil data dari map
	fmt.Println(person["bulan"])

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Azzin"
	book["wrong"] = "ups"

	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(book["ttile"])
	fmt.Println(book["author"])

}
