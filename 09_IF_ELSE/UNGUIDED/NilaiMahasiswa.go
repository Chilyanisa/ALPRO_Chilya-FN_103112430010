package main

import (
    "fmt"
)

func main() {
    var nilai float64
    fmt.Print("nilai mahasiswa: ")
    fmt.Scan(&nilai)

    if nilai > 90 {
        fmt.Println("A")
    } else if nilai >= 80 {
        fmt.Println("AB")
    } else if nilai >= 70 {
        fmt.Println("B")
    } else {
        fmt.Println("C")
    }
}