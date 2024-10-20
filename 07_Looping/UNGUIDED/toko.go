package main

import "fmt"

func main() {
    var jumlahBarang, totalPoin int
    fmt.Print("Masukkan jumlah barang yang dibeli: ")
    fmt.Scan(&jumlahBarang)

    // Setiap barang memberi 10 poin
    for i := 1; i <= jumlahBarang; i++ {
        totalPoin += 10
        // Tambahan 5 poin jika barang lebih dari 5
        if i > 5 {
            totalPoin += 5
        }
    }

    fmt.Printf("Total poin yang didapatkan: %d\n", totalPoin)
}