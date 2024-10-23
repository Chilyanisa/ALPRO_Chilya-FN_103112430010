package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Print("Masukkan nilai N: ")
    fmt.Scanln(&N)

    fmt.Println("Kuadrat dari bilangan 1 sampai", N, ":")
    for i := 1; i <= N; i++ {
        kuadrat := i * i
        fmt.Println(kuadrat)
    }
}