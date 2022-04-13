package main

import "fmt"

func endApp() {
	message := recover() // menangkap value error di panic
	if message != nil {
		fmt.Println("Aplikasi error dgn message:", message)
	}

	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(false)
}
