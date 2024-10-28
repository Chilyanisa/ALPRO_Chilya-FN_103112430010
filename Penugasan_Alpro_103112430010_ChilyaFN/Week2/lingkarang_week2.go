package main

import (
        "fmt"
        "math"
)

func main() {
        var jariJari float64

        fmt.Print("Masukkan jari-jari lingkaran: ")
        fmt.Scanln(&jariJari)

        luas := math.Pi * math.Pow(jariJari, 2)
        keliling := 2 * math.Pi * jariJari

        fmt.Printf("Luas lingkaran: %.2f\n", luas)
        fmt.Printf("Keliling lingkaran: %.2f\n", keliling)
}