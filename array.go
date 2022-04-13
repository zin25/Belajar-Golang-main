package main

import "fmt"

func main() {

	var nama [3]string

	nama[0] = "Azzin"
	nama[1] = "Maharil"
	nama[2] = "Ganteng"

	fmt.Println(nama[0], nama[1], nama[2])

	var age = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(age)
	fmt.Println(len(age))

}
