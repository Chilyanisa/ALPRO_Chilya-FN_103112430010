package main

import "fmt"

func main() {
        var bilangan int

        fmt.Print("Masukkan bilangan tiga digit: ")
        fmt.Scan(&bilangan)

        ratusan := bilangan / 100
        puluhan := (bilangan % 100) / 10
        satuan := bilangan % 10

        if ratusan > puluhan && puluhan > satuan {
                fmt.Println("Bilangan terurut mengecil")
        } else {
                fmt.Println("Bilangan tidak terurut mengecil")
        }
}