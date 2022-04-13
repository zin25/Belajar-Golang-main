package main

import (
	"fmt"
)

func main() {

	for umur := 1; umur <= 10; umur++ {
		fmt.Println("perulangan ke :", umur)
	}

	slice := []string{"Azzin", "Feby", "Ridho", "Dhody"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "azzin"
	person["title"] = "prigrimmir"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
