package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak bintang pada setiap baris
func printStars(n, i int) {
	if i > n {
		return
	}
	fmt.Println(starString(i))
	printStars(n, i+1)
}

// Fungsi rekursif untuk membuat string bintang
func starString(k int) string {
	if k == 0 {
		return ""
	}
	return "*" + starString(k-1)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)

	printStars(n, 1)
}