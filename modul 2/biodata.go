package main

import "fmt"

func main() {
    var (
        nama string
        nim  string
        kelas string
		prodi string
    )

    // Meminta input dari pengguna
    fmt.Print("Masukkan Nama: ")
    fmt.Scanln(&nama)
    
    fmt.Print("Masukkan NIM: ")
    fmt.Scanln(&nim)
    
    fmt.Print("Masukkan Kelas: ")
    fmt.Scanln(&kelas)

	fmt.Print("Masukkan Prodi: ")
    fmt.Scanln(&prodi)

    // Menampilkan resume singkat
    fmt.Println("\nBiodata Mahasiswa:")
    fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s.\n", nama, kelas, nim, prodi)
}