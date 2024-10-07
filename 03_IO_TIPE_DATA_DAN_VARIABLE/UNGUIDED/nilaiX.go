package main

import (
	"fmt"
)

func main() {
	var x int

	// Meminta input nilai f(x) dari pengguna
	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scanln(&x)

	// Menghitung nilai x berdasarkan persamaan f(x) = (x / (x + 5)) + 5
	if x == -5 {
		fmt.Println("Nilai x tidak boleh -5, karena menyebabkan pembagian dengan nol.")
	} else {
		result := (2.0 /  float64(x+5)) + 5.0
		fmt.Printf("Hasil dari f(x) = (2 / (x + 5)) + 5 adalah: %.2f\n", result)
	}
}