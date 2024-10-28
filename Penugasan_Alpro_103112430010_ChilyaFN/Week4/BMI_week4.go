package main

import (
        "fmt"
        "math"
)

type Person struct {
        name    string
        weight  float64
        height  float64
        bmi     float64
}

func main() {
        var person Person

        fmt.Print("Nama: ")
        fmt.Scanln(&person.name)

        fmt.Print("Berat (kg): ")
        fmt.Scanln(&person.weight)

        fmt.Print("Tinggi (m): ")
        fmt.Scanln(&person.height)

        if person.height == 0 {
                fmt.Println("Tinggi badan tidak boleh nol")
                return
        }

        person.bmi = person.weight / math.Pow(person.height, 2)

        fmt.Printf("BMI: %.2f\n", person.bmi)
}