package main

import "fmt"

func main() {
    var umur int
    var kewarganegaraan string

    fmt.Print("masukkan umur anda: ")
    fmt.Scan(&umur)

    fmt.Print("masukkan kewarganegaraan anda (contoh: WNI/WNA): ")
    fmt.Scan(&kewarganegaraan)

    if umur >= 17 && (kewarganegaraan == "WNI" || kewarganegaraan == "wni") {
        fmt.Println("anda bisa mengikuti pemilu")
    } else {
        fmt.Println("anda tidak bisa mengikuti pemilu karena")
        if umur < 17 {
            fmt.Println("umur anda belum mencapai 17 tahun")
        }
        if kewarganegaraan != "WNI" && kewarganegaraan != "wni" {
            fmt.Println("anda bukan WNI")
        }
    }
}