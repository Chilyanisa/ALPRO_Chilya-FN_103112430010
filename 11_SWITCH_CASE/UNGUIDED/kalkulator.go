package main

import "fmt"

func main() {
        var jenisKendaraan string
        var durasiParkir int

        fmt.Print("Masukkan jenis kendaraan (motor, mobil, truk): ")
        fmt.Scanln(&jenisKendaraan)

        fmt.Print("Masukkan durasi parkir (dalam jam): ")
        fmt.Scanln(&durasiParkir)

        if durasiParkir < 1 {
                durasiParkir = 1
        }

        var tarifPerJam int
        switch jenisKendaraan {
        case "motor":
                tarifPerJam = 2000
        case "mobil":
                tarifPerJam = 5000
        case "truk":
                tarifPerJam = 8000
        default:
                fmt.Println("Jenis kendaraan tidak valid")
                return
        }

        totalBiaya := tarifPerJam * durasiParkir
        fmt.Printf("Total biaya parkir: Rp %d\n", totalBiaya)
}