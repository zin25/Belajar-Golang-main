package main

import "fmt"

//ciri variadic adalah ... titik tiga setelah argument
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 11, 11, 11, 11, 11)
	fmt.Println(total)

	//slice
	slice := []int{10, 10, 10, 01, 01, 100}
	total = sumAll(slice...)
	fmt.Println(total)
}
