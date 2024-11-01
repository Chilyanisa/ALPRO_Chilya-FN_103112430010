package main

import (
    "fmt"
)

func main() {
    var bilangan int
    fmt.Print("Masukkan suatu bilangan bulat: ")
    fmt.Scan(&bilangan)

    if bilangan < 0 && bilangan%2 == 0 {
        fmt.Println("Bilangan adalah: genap negatif")
    } else {
        fmt.Println("Bilangan adalah: bukan")
    }
}
