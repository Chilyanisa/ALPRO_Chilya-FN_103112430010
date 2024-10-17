package main

import "fmt"

// Fungsi untuk menghitung jumlah bilangan dari 1 hingga n
func sum(n int) int {
    if n == 1 {
        return 1
    }
    return n + sum(n-1) // Rekursi untuk menjumlahkan bilangan
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan bulat positif n: ")
    fmt.Scan(&n) // Input dari pengguna

    hasil := sum(n) // Memanggil fungsi sum
    fmt.Printf("Jumlah deret angka dari 1 hingga %d adalah: %d\n", n, hasil)
}