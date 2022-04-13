package main

import "fmt"

func getBio() (firstName, lastName string, age int) {
	firstName = "Azin"
	lastName = "Maharil"
	age = 22

	return
}

func main() {
	a, b, c := getBio()
	fmt.Println(a, b)
	fmt.Println(c)
}
