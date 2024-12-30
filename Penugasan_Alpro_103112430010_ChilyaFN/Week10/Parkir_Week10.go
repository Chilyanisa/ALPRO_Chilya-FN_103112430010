package main

import (
	"fmt"
)

func main() {
	// Input waktu parkir (jam dan menit)
	var h1, m1, h2, m2 int
	fmt.Println("Masukkan jam dan menit mulai parkir (h1 m1) dan selesai parkir (h2 m2):")
	fmt.Scan(&h1, &m1, &h2, &m2)

	// Menghitung total menit untuk waktu mulai dan selesai
	startMinutes := h1*60 + m1
	endMinutes := h2*60 + m2

	// Menghitung durasi parkir
	durationMinutes := endMinutes - startMinutes

	// Mengubah durasi ke dalam jam dan menit
	hh := durationMinutes / 60
	mm := durationMinutes % 60

	// Menampilkan hasil
	fmt.Printf("Durasi parkir: %d jam %d menit\n", hh, mm)
}

//PSEUDOCODE
//Input Data:
//Masukkan nilai h1 dan m1 untuk jam dan menit mulai parkir.
//Masukkan nilai h2 dan m2 untuk jam dan menit selesai parkir.

//Konversi ke Menit:
//Hitung startMinutes = (h1 * 60) + m1.
//Hitung endMinutes = (h2 * 60) + m2.

//Hitung Durasi:
//Durasi dalam menit = endMinutes - startMinutes.

//Konversi Durasi ke Jam dan Menit:
//Jam (hh) = durationMinutes / 60.
//Menit (mm) = durationMinutes % 60.

//Tampilkan Hasil:
//Cetak hh dan mm sebagai durasi parkir.