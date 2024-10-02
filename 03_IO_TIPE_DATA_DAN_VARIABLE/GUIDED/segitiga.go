package main

import (
	"fmt"
)

func main() {
	var alas, tinggi int

	fmt.Print("Masukkan panjang alas segitiga: ")
	fmt.Scanln(&alas)

	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scanln(&tinggi)

	luas := (alas * tinggi) / 2

	fmt.Printf("Luas segitiga dengan alas %d dan tinggi %d adalah %d\n", alas, tinggi, luas)
}
