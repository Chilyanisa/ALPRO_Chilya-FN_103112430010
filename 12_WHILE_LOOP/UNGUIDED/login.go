package main

import "fmt"

func main() {
    const maxAttempts = 3
    var password string
    var correctPassword = "chilya05" // ganti dengan password yang diinginkan

    for i := 1; i <= maxAttempts; i++ {
        fmt.Print("Masukkan password: ")
        fmt.Scanln(&password)

        if password == correctPassword {
            fmt.Println("Login berhasil!")
            return
        } else {
            fmt.Println("Password salah. Silakan coba lagi.")
        }
    }

    fmt.Println("Login ditolak. Terlalu banyak percobaan yang gagal.")
}