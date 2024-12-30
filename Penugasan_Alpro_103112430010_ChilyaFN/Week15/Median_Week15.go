package main

import (
	"fmt"
)

func main() {
	var nilaiY int
	fmt.Print("Masukkan nilai Y: ")
	fmt.Scan(&nilaiY)

	bilangan := make([]int, 9)
	fmt.Println("Masukkan 9 bilangan (0 atau", nilaiY, "):")
	for i := 0; i < 9; i++ {
		fmt.Scan(&bilangan[i])
		if bilangan[i] != 0 && bilangan[i] != nilaiY {
			fmt.Println("Bilangan harus bernilai 0 atau", nilaiY)
			return
		}
	}

	// Mengurutkan bilangan secara manual
	for i := 0; i < len(bilangan)-1; i++ {
		for j := i + 1; j < len(bilangan); j++ {
			if bilangan[i] > bilangan[j] {
				bilangan[i], bilangan[j] = bilangan[j], bilangan[i]
			}
		}
	}

	// Median adalah elemen di posisi tengah (index ke-4 dari array yang diurutkan)
	median := bilangan[4]

	fmt.Printf("Median bernilai %d\n", median)
}

//PSEUDECODE
//program Median
//kamus
//j, y, bilangan, total : integer
//algoritma
//input(y)
//total = 0
//for j = 1 to 9 do
//input(bilangan)
//total = total + bilangan
//end for
//if total >= y * 5 then
//output("Median bernilai", y)
//else
//output("Median bernilai 0")
//end if
//endprogram