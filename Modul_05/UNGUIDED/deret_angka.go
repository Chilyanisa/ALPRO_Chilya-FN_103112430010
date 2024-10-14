package main

import (
	"fmt"
)

// Fungsi rekursif untuk menjumlahkan angka dari 1 hingga n
func sum(n int) int {
	if n == 1 {
		return 1
	}
	return n + sum(n-1)
}

func main() {
	var n int

	// Meminta input dari user
	fmt.Print("Masukkan bilangan bulat positif n: ")
	fmt.Scan(&n)

	// Memanggil fungsi rekursif dan menampilkan hasilnya
	fmt.Printf("Jumlah deret angka dari 1 hingga %d adalah: %d\n", n, sum(n))
}