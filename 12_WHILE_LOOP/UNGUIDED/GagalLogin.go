package main

import (
	"fmt"
)

func main() {
	const correctUsername = "Admin"
	const correctPassword = "Admin"

	var failedAttempts int

	for {
		var username, password string
		fmt.Print("Masukkan username: ")
		fmt.Scanln(&username)

		fmt.Print("Masukkan password: ")
		fmt.Scanln(&password)

		if username == correctUsername && password == correctPassword {
			break
		} else {
			fmt.Println("Username atau password salah. Coba lagi.")
			failedAttempts++
		}
	}

	fmt.Printf("%d percobaan gagal login\n", failedAttempts)
}