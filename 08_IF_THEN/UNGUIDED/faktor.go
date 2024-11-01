package main

import (
    "fmt"
)

func main() {
    var x, y int
    fmt.Print("Masukkan bilangan bulat positif x: ")
    fmt.Scan(&x)
    fmt.Print("Masukkan bilangan bulat positif y: ")
    fmt.Scan(&y)

    if y % x == 0 {
        fmt.Println("x adalah faktor dari y: true")
    } else {
        fmt.Println("x adalah faktor dari y: false")
    }

    if x % y == 0 {
        fmt.Println("y adalah faktor dari x: true")
    } else {
        fmt.Println("y adalah faktor dari x: false")
    }
}
