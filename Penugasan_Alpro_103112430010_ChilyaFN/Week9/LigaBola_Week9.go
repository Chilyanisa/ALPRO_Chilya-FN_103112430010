package main

import (
	"fmt"
)

func main() {
	// Masukan jumlah gol dari empat tim
	var a, b, c, d int
	fmt.Println("Masukkan empat jumlah gol:")
	fmt.Scan(&a, &b, &c, &d)

	// Menentukan jumlah gol terbanyak dan tersedikit
	max := a
	min := a

	// Bandingkan untuk menemukan nilai maksimum
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if d > max {
		max = d
	}

	// Bandingkan untuk menemukan nilai minimum
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	if d < min {
		min = d
	}

	// Menampilkan hasil
	fmt.Printf("Jumlah gol terbanyak: %d\n", max)
	fmt.Printf("Jumlah gol tersedikit: %d\n", min)
}

//PSEUDOCODE

//Masukkan Data:
//Input empat angka (misalkan a, b, c, d) yang mewakili jumlah gol.

//Inisialisasi:
//Tetapkan max = a.
//Tetapkan min = a.

//Cari Nilai Maksimum:
//Jika b > max, maka max = b.
//Jika c > max, maka max = c.
//Jika d > max, maka max = d.

//Cari Nilai Minimum:
//Jika b < min, maka min = b.
//Jika c < min, maka min = c.
//Jika d < min, maka min = d.

//Tampilkan Hasil:
//Cetak nilai max sebagai jumlah gol terbanyak.
//Cetak nilai min sebagai jumlah gol tersedikit.