package main

import "fmt"

// Fungsi untuk menghitung faktorial secara iteratif
func faktorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var n int

	// Meminta input bilangan dari pengguna
	fmt.Print("Masukkan bilangan bulat non-negatif: ")
	fmt.Scan(&n)

	// Menghitung dan menampilkan hasil faktorial
	if n >= 0 {
		fmt.Printf("Faktorial dari %d adalah: %d\n", n, faktorial(n))
	} else {
		fmt.Println("Masukkan bilangan non-negatif!")
	}
}