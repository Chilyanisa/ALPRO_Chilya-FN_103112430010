package main

import (
	"fmt"
	"math"
)

func main() {
	var k int

	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	result := 1.0

	for i := 1; i <= k; i++ {
		numerator := math.Pow(float64(4*k+2), 2)
		denominator := float64((4*k +1) * (4*k + 3)) 
		result *= numerator / denominator
	}

	fmt.Printf("Hasil dari operasi fungsi = %.10f\n", result)
}