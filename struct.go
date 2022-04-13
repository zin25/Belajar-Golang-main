package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHii(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var azzin Customer

	azzin.Name = "Azzin"
	azzin.Address = "Indonesia"
	azzin.Age = 23

	azzin.sayHii("Asep")

	// // fmt.Println(azzin)
	// fmt.Println(azzin.Name)
	// fmt.Println(azzin.Address)
	// fmt.Println(azzin.Age)

	// asep := Customer{
	// 	Name:    "Asep",
	// 	Address: "Jawa Timur",
	// 	Age:     25,
	// }

	// fmt.Println(asep)

}
