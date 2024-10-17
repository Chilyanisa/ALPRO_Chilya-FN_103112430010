package main

import (
	"fmt"
	"strconv"
)

// Fungsi untuk menghitung luas segitiga
func hitungLuas(alas, tinggi float64) float64 {
	return (alas * tinggi) / 2
}

// Fungsi untuk menghilangkan nol di belakang koma
func formatLuas(luas float64) string {
	return strconv.FormatFloat(luas, 'f', -1, 64)
}

func main() {
	var n int

	// Meminta input jumlah segitiga
	fmt.Print("Masukkan jumlah segitiga (n): ")
	fmt.Scan(&n)

	// Memproses input alas dan tinggi untuk masing-masing segitiga
	for i := 1; i <= n; i++ {
		var alas, tinggi float64
		fmt.Printf("Masukkan alas segitiga %d: ", i)
		fmt.Scan(&alas)
		fmt.Printf("Masukkan tinggi segitiga %d: ", i)
		fmt.Scan(&tinggi)

		// Menghitung dan menampilkan luas segitiga
		luas := hitungLuas(alas, tinggi)
		fmt.Printf("Luas segitiga %d: %s\n", i, formatLuas(luas))
	}
}