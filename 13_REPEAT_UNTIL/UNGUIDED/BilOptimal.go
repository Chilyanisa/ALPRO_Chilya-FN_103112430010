package main

import "fmt"

func main() {
	var number float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scanln(&number)

	if number <= 0 {
		fmt.Println("Input harus bilangan desimal positif.")
		return
	}

	fmt.Println("Hasil perhitungan tiap perulangan:")

	rounded := int(number)
	if float64(rounded) < number {
		rounded++ 
	}

	for sum := number; sum < float64(rounded); sum += 0.1 {
		fmt.Printf("%.1f\n", sum)
	}
}