package main

import (
    "fmt"
)

func main() {
    var jumlahOrang int
    fmt.Print("Masukkan jumlah orang yang akan melakukan touring: ")
    fmt.Scan(&jumlahOrang)

    jumlahMotor := (jumlahOrang + 1) / 2
    fmt.Printf("Jumlah motor yang diperlukan: %d\n", jumlahMotor)
}
