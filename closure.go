//berinteraksi dengan data yang ada disekitarnya

package main

import "fmt"

func main() {
	name := "Azzin"
	counter := 0

	increment := func() {
		name = "maharil"
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
