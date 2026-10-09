package main
import "fmt"

func main() {
	var harga, diskon, total float64
	fmt.Scan(&harga, &diskon)
	total = harga - (harga * diskon / 100)
	fmt.Printf("%.0f\n", total)
}