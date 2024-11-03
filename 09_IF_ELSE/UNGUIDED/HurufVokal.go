package main

import "fmt"

func main() {
    var huruf rune
    fmt.Print("masukkan huruf: ")
    fmt.Scanf("%c", &huruf)

    if huruf == 'A' || huruf == 'I' || huruf == 'U' || huruf == 'E' || huruf == 'O' ||
       huruf == 'a' || huruf == 'i' || huruf == 'u' || huruf == 'e' || huruf == 'o' {
        fmt.Println("huruf vokal")
    } else {
        fmt.Println("huruf konsonan")
    }
}
