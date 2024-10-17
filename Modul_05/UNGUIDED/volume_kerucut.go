package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung volume kerucut
func volumeKerucut(r, t float64) float64 {
	return (1.0 / 3.0) * math.Pi * r * r * t
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n) // Input jumlah kerucut

	// Array untuk menyimpan jari-jari dan tinggi kerucut
	radii := make([]float64, n)
	heights := make([]float64, n)

	// Input jari-jari kerucut
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan jari-jari kerucut ke-%d: ", i+1)
		fmt.Scan(&radii[i])
	}

	// Input tinggi kerucut
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan tinggi kerucut ke-%d: ", i+1)
		fmt.Scan(&heights[i])
	}

	// Menghitung dan menampilkan volume tiap kerucut
	for i := 0; i < n; i++ {
		volume := volumeKerucut(radii[i], heights[i])
		fmt.Printf("Volume kerucut ke-%d adalah: %.2f\n", i+1, volume)
	}
}