package main
import (
	"fmt"
)

func main() {
	var age int
	fmt.Print("Masukkan usia Anda: ")
	fmt.Scan(&age)

	switch {
	case age >= 0 && age <= 12:
		fmt.Println("Kategori: Anak-anak")
	case age >= 13 && age <= 17:
		fmt.Println("Kategori: Remaja")
	case age >= 18 && age <= 64:
		fmt.Println("Kategori: Dewasa")
	case age >= 65:
		fmt.Println("Kategori: Lansia")
	default:
		fmt.Println("Usia tidak valid")
	}
}