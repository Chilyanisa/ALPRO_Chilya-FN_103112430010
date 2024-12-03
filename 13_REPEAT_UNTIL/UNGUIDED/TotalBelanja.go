package main

import (
	"fmt"
)

func main() {
	var total int = 0
	var harga int

	for {
		fmt.Print("Masukkan harga barang (ketik 0 untuk selesai): ")
		_, err := fmt.Scan(&harga) 

		if err != nil {
			fmt.Println("Masukkan angka yang valid.")
			continue
		}

		if harga == 0 {
			break 
		}

		if harga > 0 {
			total += harga 
		} else {
			fmt.Println("Harga tidak valid. Masukkan angka positif.")
		}
	}

	fmt.Printf("Total belanja Anda: %d\n", total)
}