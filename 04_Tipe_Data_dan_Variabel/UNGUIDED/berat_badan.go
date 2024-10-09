package main

import (
    "fmt"
    "math"
)

func calculateWeight(bmi, height float64) float64 {
    return bmi * (height * height)
}

func main() {
    var bmi, height float64

    fmt.Print("Masukkan nilai BMI: ")
    fmt.Scan(&bmi)
    fmt.Print("Masukkan tinggi badan (meter): ")
    fmt.Scan(&height)

    weight := calculateWeight(bmi, height)
    fmt.Printf("Berat badan Anda adalah: %d kg\n", int(math.Round(weight)))
}