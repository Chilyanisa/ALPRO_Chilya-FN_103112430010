package main

import (
	"fmt"
)

func main() {
	var x float64
	var hasil float64

	fmt. Println("Masukan nilai x: ")
	fmt. Scanln(&x)

	hasil = (2/ (x+5)) +5
	fmt.Printf("Hasil dari f(x)= %.0f adalah %.1f\n ", x, hasil)
}