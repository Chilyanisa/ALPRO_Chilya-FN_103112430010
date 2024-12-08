package main

import (
	"fmt"
)

func isPerfectNumber(num int) bool {
	if num <= 0 {
		return false
	}

	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	return sum == num
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&n)

	if isPerfectNumber(n) {
		fmt.Printf("Ya, %d adalah bilangan sempurna.\n", n)
	} else {
		fmt.Printf("Tidak, %d bukan bilangan sempurna.\n", n)
	}
}