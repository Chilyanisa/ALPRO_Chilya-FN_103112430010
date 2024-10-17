package main

import "fmt"

// Fungsi untuk menghitung pangkat dengan menggunakan perkalian dan perulangan
func pangkat(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func main() {
	var base, exponent int

	// Meminta input bilangan dasar dan pangkat
	fmt.Print("Masukkan bilangan pertama (basis): ")
	fmt.Scan(&base)
	fmt.Print("Masukkan bilangan kedua (pangkat): ")
	fmt.Scan(&exponent)

	// Menghitung hasil pemangkatan dan menampilkan hasilnya
	result := pangkat(base, exponent)
	fmt.Printf("%d pangkat %d adalah: %d\n", base, exponent, result)
}