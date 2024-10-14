package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Menginisialisasi generator angka acak
	target := rand.Intn(100) + 1     // Memilih angka acak antara 1 hingga 100

	fmt.Println("Tebak angka antara 1 hingga 100. Anda punya 5 kali kesempatan!")
	playGame(target, 1) // Memulai permainan dengan percobaan pertama
}

func playGame(target, attempt int) {
	var guess int

	// Jika percobaan melebihi 5, permainan berakhir
	if attempt > 5 {
		fmt.Printf("Kesempatan habis! Angka yang benar adalah %d\n", target)
		return
	}

	fmt.Printf("Percobaan %d: Masukkan tebakan Anda: ", attempt)
	fmt.Scan(&guess)

	// Cek apakah tebakan benar, terlalu besar, atau kecil
	if guess == target {
		fmt.Println("Selamat! Anda berhasil menebak angka dengan benar.")
	} else if guess < target {
		fmt.Println("Terlalu kecil!")
		playGame(target, attempt+1)
	} else {
		fmt.Println("Terlalu besar!")
		playGame(target, attempt+1)
	}
}