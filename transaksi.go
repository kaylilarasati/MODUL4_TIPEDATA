package main

import "fmt"

type DataTransaksi struct {
	namaBarang   string
	jumlahBarang int
	hargaBarang  float64
	totalHarga   float64
}

func main() {
	var data DataTransaksi
	fmt.Scan(&data.namaBarang, &data.jumlahBarang, &data.hargaBarang)
	data.totalHarga = float64(data.jumlahBarang) * data.hargaBarang
	fmt.Printf("Nama Barang: %s\n", data.namaBarang)
	fmt.Printf("Jumlah Barang: %d\n", data.jumlahBarang)
	fmt.Printf("Harga Barang: %.2f\n", data.hargaBarang)
	fmt.Printf("Total Harga: Rp %.2f\n", data.totalHarga)
}
