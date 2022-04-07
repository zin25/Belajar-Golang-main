package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye" + name
}

func main() {
	//memasukan function ke variable

	sayGoodBye := getGoodBye

	result := sayGoodBye("Azzin")
	fmt.Println(result)
}
