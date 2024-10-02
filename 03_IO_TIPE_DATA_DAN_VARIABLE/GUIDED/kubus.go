package main

import (
	"fmt"
)

func main() {
	var sisi int

	fmt.Print("Masukkan panjang sisi kubus: ")
	fmt.Scanln(&sisi)

	volume := sisi * sisi * sisi

	fmt.Printf("Volume kubus dengan sisi %d adalah %d\n", sisi, volume)
}
