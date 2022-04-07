package main

import "fmt"

func getFullName() (string, string) {
	return "Azzin", "Maharil"
}

//if u want to ignore u can use "_" example, you don't have last name, so just change with "_"
func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
	//fmt.Println(lastName)
}
