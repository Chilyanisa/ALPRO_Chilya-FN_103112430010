package main

import "fmt"

func main() {
    var pH float64

    fmt.Print("Masukkan nilai pH: ")
    fmt.Scan(&pH)

    if pH >= 0 && pH <= 14 {
        if pH >= 6.5 && pH <= 8.6 {
            fmt.Println("Air layak minum")
        } else {
            fmt.Println("Air tidak layak minum")
        }
    } else {
        fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
    }
}