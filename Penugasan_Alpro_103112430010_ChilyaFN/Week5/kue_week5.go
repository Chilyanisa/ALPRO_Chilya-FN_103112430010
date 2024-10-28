package main

import "fmt"

func main() {
        var n int
        fmt.Scan(&n)

        var totalTepung int
        for i := 0; i < n; i++ {
                var tepungPerKue int
                fmt.Scan(&tepungPerKue)

                fmt.Printf("Untuk jenis kue ke-%d, dibutuhkan %d gram tepung.\n", i+1, tepungPerKue*10)

                totalTepung += tepungPerKue * 10
        }

        fmt.Printf("Total tepung yang dibutuhkan adalah %d gram.\n", totalTepung)
}