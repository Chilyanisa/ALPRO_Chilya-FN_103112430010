package main

import (
    "fmt"
)

func main() {
    var jamKerjaMinggu, upahPerJam, jamNormal, lembur, totalGaji float64
    const jamNormalMinggu = 40

    // Input jam kerja per minggu dan upah per jam
    fmt.Print("Masukkan jumlah jam kerja per minggu: ")
    fmt.Scan(&jamKerjaMinggu)
    fmt.Print("Masukkan upah per jam: ")
    fmt.Scan(&upahPerJam)

    // Hitung lembur jika jam kerja lebih dari 40 jam
    if jamKerjaMinggu > jamNormalMinggu {
        lembur = jamKerjaMinggu - jamNormalMinggu
        jamNormal = jamNormalMinggu
    } else {
        jamNormal = jamKerjaMinggu
        lembur = 0
    }

    // Hitung total gaji bulanan
    totalGaji = (jamNormal * upahPerJam * 4) + (lembur * 1.5 * upahPerJam * 4)

    // Tampilkan cara menghitung dan hasil
    fmt.Printf("Perhitungan:\n")
    fmt.Printf("Jam Normal: %.2f jam\n", jamNormal)
    fmt.Printf("Lembur: %.2f jam\n", lembur)
    fmt.Printf("Gaji dari jam normal: %.2f x %.2f x 4 = %.2f\n", jamNormal, upahPerJam, jamNormal*upahPerJam*4)
    fmt.Printf("Gaji dari lembur: %.2f x 1.5 x %.2f x 4 = %.2f\n", lembur, upahPerJam, lembur*1.5*upahPerJam*4)
    fmt.Printf("Total gaji bulanan: %.2f\n", totalGaji)
}
