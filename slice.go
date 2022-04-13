package main

import "fmt"

func main() {

	var bulan = [...]string{

		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = bulan[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[0] = "Mei Baru"
	fmt.Println(slice1)

	var slice2 = bulan[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Azin")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(bulan)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Azzin"
	newSlice[1] = "Maharil"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, 1, cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3}

	fmt.Println(iniArray, iniSlice)

}
