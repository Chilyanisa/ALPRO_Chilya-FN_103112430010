package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&number)

	if number <= 0 {
		fmt.Println("Input harus bilangan bulat positif.")
		return
	}

	count := 0

	for number > 0 {
		number /= 10
		count++
	}

	fmt.Printf("Jumlah digit: %d\n", count)
}