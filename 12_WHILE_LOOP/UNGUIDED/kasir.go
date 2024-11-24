package main

import "fmt"

func main() {
    var totalHarga float64
    var namaBarang string
    var hargaBarang float64

    for {
        fmt.Print("Masukkan nama barang (ketik 'selesai' untuk mengakhiri): ")
        fmt.Scanln(&namaBarang)

        if namaBarang == "selesai" {
            break
        }

        fmt.Print("Masukkan harga barang: ")
        fmt.Scanln(&hargaBarang)

        totalHarga += hargaBarang

        fmt.Println("Barang", namaBarang, "berhasil ditambahkan.")
        fmt.Println("Total belanja saat ini:", totalHarga)
    }

    fmt.Println("\nTotal belanja keseluruhan:", totalHarga)
}