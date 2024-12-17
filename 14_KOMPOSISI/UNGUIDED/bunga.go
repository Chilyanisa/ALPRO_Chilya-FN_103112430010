package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var bunga, pita string
	var bungaList []string

	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	if n > 0 {
		for i := 1; i <= n; i++ {
			fmt.Printf("Bunga %d: ", i)
			fmt.Scan(&bunga)
			bungaList = append(bungaList, bunga)
		}
		pita = strings.Join(bungaList, " - ")
		fmt.Println("Pita:", pita)
	} else {
		fmt.Println("N tidak boleh nol atau negatif.")
	}

	bungaList = nil
	fmt.Println("\nMasukkan bunga, ketik 'SELESAI' untuk berhenti:")
	for i := 1; ; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)

		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}
		bungaList = append(bungaList, bunga)
	}

	pita = strings.Join(bungaList, " - ")
	fmt.Println("Pita:", pita)
	fmt.Println("Banyaknya bunga:", len(bungaList))
}