package main

import (
    "fmt"
)

func main() {
    var radius float64
    const pi = 22.0 / 7.0

    // Input jari-jari lingkaran
    fmt.Print("Masukkan jari-jari lingkaran: ")
    fmt.Scan(&radius)

    // Hitung luas lingkaran
    luas := pi * radius * radius

    // Hitung keliling lingkaran
    keliling := 2 * pi * radius

    // Tampilkan hasil tanpa nol di belakang
    fmt.Printf("Luas lingkaran: %g\n", luas)
    fmt.Printf("Keliling lingkaran: %g\n", keliling)
}