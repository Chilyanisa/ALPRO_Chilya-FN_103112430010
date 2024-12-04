package main

import (
	"fmt"
)

func main() {
	var target, donasi, totalDonasi, jumlahDonatur int

	fmt.Print("Masukkan target donasi: ")
	fmt.Scanln(&target)

	for totalDonasi < target {
		fmt.Printf("Masukkan donasi dari Donatur %d: ", jumlahDonatur+1)
		fmt.Scanln(&donasi)
		jumlahDonatur++
		totalDonasi += donasi
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, totalDonasi)
	}

	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonasi, jumlahDonatur)
}