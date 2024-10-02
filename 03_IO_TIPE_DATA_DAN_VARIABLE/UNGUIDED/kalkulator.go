package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&num1)

	fmt.Print("Masukkan operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&num2)

	
	switch operator {
	case "+":
		fmt.Printf("Hasil: %g + %g = %g\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("Hasil: %g - %g = %g\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("Hasil: %g * %g = %g\n", num1, num2, num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Hasil: %g / %g = %g\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Kesalahan: Pembagian dengan nol tidak diizinkan!")
		}
	default:
		fmt.Println("Operator tidak valid!")
	}
}