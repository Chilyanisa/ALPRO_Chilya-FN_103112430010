package main

import "fmt"

func convertTime(seconds int) (int, int, int) {
    hours := seconds / 3600
    minutes := (seconds % 3600) / 60
    secs := seconds % 60
    return hours, minutes, secs
}

func main() {
    var seconds int
    fmt.Print("Masukkan jumlah detik: ")
    fmt.Scan(&seconds)

    hours, minutes, secs := convertTime(seconds)
    fmt.Printf("Hasil konversi: %d jam %d menit %d detik\n", hours, minutes, secs)
}