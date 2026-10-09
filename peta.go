package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik (Pythagoras)
func hitungJarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func main() {
	var xa, ya float64
	var xb, yb float64
	var xc, yc float64

	// Membaca masukan koordinat titik A, B, dan C
	fmt.Scanln(&xa, &ya)
	fmt.Scanln(&xb, &yb)
	fmt.Scanln(&xc, &yc)

	// Menghitung panjang masing-masing sisi
	sisiAB := hitungJarak(xa, ya, xb, yb)
	sisiBC := hitungJarak(xb, yb, xc, yc)
	sisiCA := hitungJarak(xc, yc, xa, ya)

	// Mencari sisi terpanjang
	sisiTerpanjang := sisiAB
	if sisiBC > sisiTerpanjang {
		sisiTerpanjang = sisiBC
	}
	if sisiCA > sisiTerpanjang {
		sisiTerpanjang = sisiCA
	}

	// Menampilkan keluaran dengan 2 angka di belakang koma
	fmt.Printf("%.2f\n", sisiTerpanjang)
}