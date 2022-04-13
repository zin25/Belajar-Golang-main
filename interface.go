package main

import "fmt"

type HashName interface {
	GetName() string
}

//function sayHello, parameter dari hashname
func sayHelloo(hashName HashName) {
	fmt.Println("Hello", hashName.GetName()) // Print Hello dan ambil parameneter hasname
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var azin Person
	azin.Name = "Azzin"

	sayHelloo(azin)

	cat := Animal{
		Name: "Push",
	}
	sayHelloo(cat)
}
