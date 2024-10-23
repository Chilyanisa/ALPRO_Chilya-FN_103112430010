package main

import (
    "fmt"
)

func main() {
    var qirat int
    fmt.Print("Masukkan nilai uang dalam satuan qirat: ")
    fmt.Scanln(&qirat)

    fals := qirat / 6
    sisaQirat := qirat % 6
    dirham := fals / 10
    sisaFals := fals % 10
    dinar := dirham / 10
    sisaDirham := dirham % 10

    fmt.Println("Hasil konversi:")
    fmt.Printf("Dinar: %d\n", dinar)
    fmt.Printf("Dirham: %d\n", sisaDirham)
    fmt.Printf("Fals: %d\n", sisaFals)
    fmt.Printf("Qirat: %d\n", sisaQirat)
}