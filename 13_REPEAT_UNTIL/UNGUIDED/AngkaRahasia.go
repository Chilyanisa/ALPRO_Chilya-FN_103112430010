package main

import (
	"fmt"
)

func main() {
	angkaRahasia := 7
	var tebakan int

	for {
		fmt.Print("Tebak angka (1-10): ")
		_, err := fmt.Scan(&tebakan)

		if err != nil {
			fmt.Println("Masukkan angka yang valid.")
			fmt.Scanln() 
			continue
		}

		if tebakan == angkaRahasia {
			fmt.Println("Selamat, tebakan Anda benar!")
			break
		} else {
			fmt.Println("Tebakan Anda salah, coba lagi.")
		}
	}
}