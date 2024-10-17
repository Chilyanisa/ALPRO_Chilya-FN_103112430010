package main

import "fmt"

// Fungsi untuk menghitung perkalian menggunakan penjumlahan berulang
func multiply(a, b int) int {
	if b == 0 {
		return 0
	}
	return a + multiply(a, b-1)
}

func main() {
	var a, b int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan kedua: ")
	fmt.Scan(&b)

	// Menampilkan hasil perkalian
	result := multiply(a, b)
	fmt.Printf("Hasil perkalian %d dan %d adalah: %d\n", a, b, result)
}