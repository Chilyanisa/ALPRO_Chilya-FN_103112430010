package main

import "fmt"

func main() {
    var beratParsels int
    var biayaTotal int

    fmt.Print("Masukkan berat parsel (dalam gram): ")
    fmt.Scan(&beratParsels)

    beratKg := beratParsels / 1000            
    sisaGram := beratParsels % 1000           

    biayaPerKg := 10000
    biayaTotal = beratKg * biayaPerKg

    if beratKg > 10 {
        fmt.Printf("Total berat: %d kg %d gram\n", beratKg, sisaGram)
        fmt.Printf("Detail biaya: Rp. %d\n", biayaTotal)
        fmt.Printf("Total biaya: Rp. %d\n", biayaTotal)
    } else {
        if sisaGram >= 500 {
            biayaTotal += sisaGram * 5  
        } else if sisaGram > 0 {
            biayaTotal += sisaGram * 15 
        }

        fmt.Printf("Total berat: %d kg %d gram\n", beratKg, sisaGram)
        fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", beratKg*biayaPerKg, biayaTotal-(beratKg*biayaPerKg))
        fmt.Printf("Total biaya: Rp. %d\n", biayaTotal)
    }
}