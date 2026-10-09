package main

import "fmt"

func main(){
	var jumlah int
	var total float64
	fmt.Scan(&jumlah)
	const harga = 20000
	const pajak = 0.11
	total = float64(jumlah) * harga + (float64(jumlah) * harga * pajak)
	fmt.Printf("%.0f\n", total)
}
