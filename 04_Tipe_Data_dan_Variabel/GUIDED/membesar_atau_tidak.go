package main

import "fmt"

func isAscending(num int) bool {
    if num < 100 || num > 999 {
        return false
    }

    hundreds := num / 100
    tens := (num / 10) % 10
    ones := num % 10

    return hundreds < tens && tens < ones
}

func main() {
    var num int
    fmt.Print("Masukkan bilangan tiga digit: ")
    fmt.Scan(&num)

    if isAscending(num) {
        fmt.Println("True")
    } else {
        fmt.Println("False")
    }
}
