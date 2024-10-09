package main

import "fmt"

func calculateTotal(price, discount int) int {
    return price - (price * discount / 100)
}

func main() {
    var price, discount int

    fmt.Print("Masukkan total belanja: ")
    fmt.Scan(&price)
    fmt.Print("Masukkan diskon (%): ")
    fmt.Scan(&discount)

    total := calculateTotal(price, discount)
    fmt.Printf("Total belanja setelah diskon: %d\n", total)
}
