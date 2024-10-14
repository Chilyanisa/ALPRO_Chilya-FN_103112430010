package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bilangan genap dari 1 hingga 50:")

	printEven(2) // Memulai dari angka genap pertama

}

// Fungsi rekursif untuk mencetak bilangan genap
func printEven(n int) {
	if n > 50 {
		return
	}
	fmt.Println(n)
	printEven(n + 2) // Melanjutkan ke bilangan genap berikutnya
}