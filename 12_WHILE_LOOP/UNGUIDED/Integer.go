package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x (x >= y): ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scanln(&y)

	if x < y || y <= 0 {
		fmt.Println("Input tidak valid. Pastikan x >= y dan y > 0.")
		return
	}

	result := 0

	for x >= y {
		x -= y
		result++
	}

	fmt.Printf("Hasil x div y: %d\n", result)
}