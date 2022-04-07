package main

import "fmt"

func main() {
	for bill := 1; bill < 50; bill++ {
		i := 0
		for bag := 1; bag < 50; bag++ {
			if bill%bag == 0 {
				i++
			}
		}
		if (i == 2) && (bill != 1) {
			fmt.Println(bill)
		}
	}
}
