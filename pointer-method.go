package main

import (
	"fmt"
)

type Man struct {
	Name string
}

func (man *Man) getName() {
	man.Name = "Mr." + man.Name
}

func main() {
	azzin := Man{"Azzin"}
	azzin.getName()

	fmt.Println(azzin.Name)
}
