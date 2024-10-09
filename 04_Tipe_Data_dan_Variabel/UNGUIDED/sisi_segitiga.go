package main

import (
    "fmt"
    "math"
    "strconv"
)

func distance(x1, y1, x2, y2 float64) float64 {
    return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func formatFloatWithoutTrailingZeros(f float64) string {
    return strconv.FormatFloat(f, 'f', -1, 64)
}

func main() {
    var xA, yA, xB, yB, xC, yC float64

    // Input koordinat titik A, B, C
    fmt.Print("Masukkan koordinat titik A (x y): ")
    fmt.Scan(&xA, &yA)
    fmt.Print("Masukkan koordinat titik B (x y): ")
    fmt.Scan(&xB, &yB)
    fmt.Print("Masukkan koordinat titik C (x y): ")
    fmt.Scan(&xC, &yC)

    // Hitung panjang sisi-sisi segitiga
    AB := distance(xA, yA, xB, yB)
    BC := distance(xB, yB, xC, yC)
    CA := distance(xC, yC, xA, yA)

    // Tentukan sisi terpanjang
    longest := math.Max(AB, math.Max(BC, CA))

    // Output tanpa trailing zero
    fmt.Printf("Sisi terpanjang segitiga adalah: %s\n", formatFloatWithoutTrailingZeros(longest))
}

