package main

import (
    "errors"
    "fmt"
)

func jumlahBilangan(x, y int) (int, error) {
    if x > y {
        return 0, errors.New("x harus lebih kecil atau sama dengan y")
    }

    total := 0
    for i := x; i <= y; i++ {
        total += i
    }
    return total, nil
}

func main() {
    var x, y int
    fmt.Print("Masukkan nilai x: ")
    fmt.Scanln(&x)
    fmt.Print("Masukkan nilai y: ")
    fmt.Scanln(&y)

    hasil, err := jumlahBilangan(x, y)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Hasil penjumlahan dari %d sampai %d adalah: %d\n", x, y, hasil)
    }
}