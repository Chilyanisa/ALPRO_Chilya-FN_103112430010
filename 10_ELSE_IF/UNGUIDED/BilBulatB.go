package main
import (
    "fmt"
)

func main() {
    var b int

    fmt.Print("Bilangan: ")
    fmt.Scan(&b)

    if b <= 1 {
        fmt.Println("Bilangan harus lebih besar dari 1.")
        return
    }

    fmt.Print("Faktor: ")
    jumlahFaktor := 0
    for i := 1; i <= b; i++ {
        if b%i == 0 {
            fmt.Print(i, " ")
            jumlahFaktor++
        }
    }
    fmt.Println()

    prima := jumlahFaktor == 2
    fmt.Println("Prima:", prima)
}