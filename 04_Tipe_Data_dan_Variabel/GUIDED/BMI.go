package main

import (
    "fmt"
    "math"
)

func calculateBMI(weight, height float64) float64 {
    return weight / math.Pow(height, 2)
}

func main() {
    var weight, height float64

    fmt.Print("Masukkan berat badan (kg): ")
    fmt.Scan(&weight)
    fmt.Print("Masukkan tinggi badan (m): ")
    fmt.Scan(&height)

    bmi := calculateBMI(weight, height)
    fmt.Printf("BMI Anda adalah: %.2f\n", bmi)
}