package main

import (
	"fmt"
)

func main() {
	var kantong1, kantong2 float64

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		if kantong1 > 9 || kantong2 > 9 {
			fmt.Println("Program berhenti karena salah satu kantong lebih dari 9 kg.")
			break
		}

		if diff := kantong1 - kantong2; diff > 9 || diff < -9 {
			fmt.Println("Program berhenti karena selisih berat kedua kantong lebih dari 9 kg.")
			break
		}

		total := kantong1 + kantong2
		if total > 150 {
			fmt.Println("Program berhenti karena total berat belanjaan melebihi 150 kg.")
			break
		}

		fmt.Println("Berat valid, lanjutkan input.")
	}

	fmt.Println("Proses selesai.")
}