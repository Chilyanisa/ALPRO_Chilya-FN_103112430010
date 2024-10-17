package main

import "fmt"

func main() {
	var a, b int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan b: ")
	fmt.Scan(&b)

	// Mencetak bilangan dari a hingga b
	for i := a; i <= b; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println() // Pindah baris setelah selesai mencetak
}