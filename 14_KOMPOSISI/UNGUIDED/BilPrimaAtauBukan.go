package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Masukkan sebuah bilangan bulat positif: ")
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan prima")
	}
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}