package main

import "fmt"

func main() {
        var number int
        fmt.Print("Masukkan bilangan: ")
        fmt.Scan(&number)

        switch {
        case number%2 != 0:
                result := number + (number + 1)
                fmt.Printf("Kategori: Bilangan Ganjil\nHasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", number, number+1, result)
        case number%2 == 0:
                result := number * (number + 1)
                fmt.Printf("Kategori: Bilangan Genap\nHasil perkalian dengan bilangan berikutnya %d * %d = %d\n", number, number+1, result)
        case number%5 == 0:
                result := number * number
                fmt.Printf("Kategori: Bilangan Kelipatan 5\nHasil kuadrat dari %d^2 = %d\n", number, result)
        case number%10 == 0:
                result := number / 10
                fmt.Printf("Kategori: Bilangan Kelipatan 10\nHasil pembagian antara %d/10 = %d\n", number, result)
        }
}