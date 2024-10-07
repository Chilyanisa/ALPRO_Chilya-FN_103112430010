package main

import (
	"fmt"
)

func main() {
	var celsius float64

	// Membaca input suhu dalam Celsius
	fmt.Print("Masukkan temperatur Celsius: ")
	fmt.Scan(&celsius)

	// Menghitung konversi ke Reamur, Fahrenheit, dan Kelvin
	reamur := celsius * 4 / 5
	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273.15

	// Menampilkan hasil konversi
	fmt.Printf("Derajat Reamur: %.2f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}