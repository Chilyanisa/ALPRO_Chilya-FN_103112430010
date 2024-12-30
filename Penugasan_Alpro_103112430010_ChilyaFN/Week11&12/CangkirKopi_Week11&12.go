package main

import (
	"fmt"
)

func main() {
	var n, m, x, y int
	fmt.Print("Masukkan n, m, x, y (dipisahkan spasi): ")
	fmt.Scanf("%d %d %d %d", &n, &m, &x, &y)

	// Hitung jumlah maksimum cangkir kopi
	jumlahGula := n / x
	jumlahKopi := m / y

	// Cangkir yang bisa dibuat adalah minimum dari jumlahGula dan jumlahKopi
	jumlahCangkir := 0
	if jumlahGula < jumlahKopi {
		jumlahCangkir = jumlahGula
	} else {
		jumlahCangkir = jumlahKopi
	}

	fmt.Printf("Jumlah cangkir kopi yang bisa dibuat: %d\n", jumlahCangkir)
}

//PSEUDECODE
//Input: n, m, x, y (bilangan bulat)
//Output: jumlahCangkir (jumlah maksimum cangkir kopi)

//Baca nilai n, m, x, dan y
//Hitung jumlah gula yang tersedia untuk membuat kopi:
   //jumlahGula = n / x
//Hitung jumlah kopi yang tersedia:
   //jumlahKopi = m / y
//Tentukan jumlah cangkir kopi yang dapat dibuat:
   //jumlahCangkir = Minimum(jumlahGula, jumlahKopi)
//Tampilkan jumlahCangkir