package main

import "fmt"

func isPalindrome(input string) bool {

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scan(&str)
	result := isPalindrome(str)
	if result == true {
		fmt.Printf("Palindrome")
	} else {
		fmt.Printf("not Palindrome")
	}
}
