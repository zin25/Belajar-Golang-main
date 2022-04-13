package main

import "fmt"

//Assertion untuk ubah type apapun semau kita

func random() interface{} {
	return "OK"
}

func main() {

	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}

}
