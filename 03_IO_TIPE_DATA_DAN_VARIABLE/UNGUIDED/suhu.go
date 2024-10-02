package main

import (
	"fmt"
)

func main() {
	var fahrenheit float64

	fmt.Print("Masukkan suhu dalam derajat Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9

	kelvin := celsius + 273.15

	fmt.Printf("Suhu dalam derajat Kelvin: %g K\n", kelvin)
}
