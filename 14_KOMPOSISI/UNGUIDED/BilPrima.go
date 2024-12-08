package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&n)

	fmt.Printf("Bilangan prima dari 1 hingga %d adalah:\n", n)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}