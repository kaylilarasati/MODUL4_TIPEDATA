package main

import "fmt"

func main() {
    var h, batas, tekanan float64

    fmt.Scan(&h, &batas)

    const p float64 = 1025.0
    const g float64 = 9.8

    tekanan = p * g * h

    fmt.Printf("%.1f\n", tekanan)
    fmt.Println(tekanan <= batas)
}