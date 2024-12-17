package main

import (
	"fmt"
)

func main() {
	var a, b, c, d string
	warna := true
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&a, &b, &c, &d)

		if a != "merah" || b != "kuning" || c != "hijau" || d != "ungu" {
			warna = false
		}
	}
	fmt.Println("BERHASIL: ", warna)
}