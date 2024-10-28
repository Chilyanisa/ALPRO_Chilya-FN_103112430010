package main

import "fmt"

func main() {
        var nilaiUjian int

        fmt.Print("Masukkan nilai ujian: ")
        fmt.Scan(&nilaiUjian)

        if nilaiUjian >= 70 {
                fmt.Println("Lulus")
        } else {
                fmt.Println("Tidak Lulus")
        }
}