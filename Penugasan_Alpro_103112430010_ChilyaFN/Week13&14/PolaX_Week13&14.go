package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("Masukkan bilangan bulat positif x: ")
	fmt.Scan(&x)

	if x <= 0 {
		fmt.Println("Masukkan bilangan bulat positif.")
		return
	}

	// Pola bagian atas
	for i := 1; i <= x; i++ {
		for j := 1; j <= x; j++ {
			if j == i || j == x-i+1 {
				fmt.Printf("%d ", i)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

//PSEUDECODE
//Input: x (bilangan bulat positif)
//Output: Pola berbentuk "X" dengan angka berdasarkan nomor baris

//Baca nilai x
//Jika x <= 0:
   //Cetak "Masukkan bilangan bulat positif."
   //Keluar dari program
//Untuk i dari 1 hingga x (bagian atas pola):
   //Untuk j dari 1 hingga x:
      //Jika j == i atau j == x-i+1:
        //Cetak nilai i di posisi tersebut
      //ii. Jika tidak:
        //Cetak spasi
   //Pindah ke baris berikutnya
//Untuk i dari x-1 hingga 1 (bagian bawah pola):
   //Untuk j dari 1 hingga x:
      //i. Jika j == i atau j == x-i+1:
        //Cetak nilai i di posisi tersebut
      //ii. Jika tidak:
        //Cetak spasi
   //Pindah ke baris berikutnya
//Program selesai
