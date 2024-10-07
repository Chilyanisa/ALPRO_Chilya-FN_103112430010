package main

import (
	"fmt"
	"math"
)

func main() {
	var r int
	const pi = 3.1415926535

	// Input jari-jari dari user
	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scan(&r)

	// Menghitung volume bola
	volume := (4.0 / 3.0) * pi * math.Pow(float64(r), 3)

	// Menghitung luas permukaan bola
	luas := 4 * pi * math.Pow(float64(r), 2)

	// Menampilkan hasil
	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}