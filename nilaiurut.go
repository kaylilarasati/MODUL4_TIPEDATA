package main
import "fmt"
func main() {
	var n, d1, d2, d3 int
	fmt.Scan(&n)

	d1 = n / 100
	d2 = n % 100 / 10
	d3 = n % 100 % 10		
	fmt.Println(d1 <= d2 && d2 <= d3)
}
// program nilaiurut
// kamus : n, d1, d2, d3 : integer
// algoritma :
// 	input(n)
// 	d1 = n / 100
// 	d2 = (n / 10) % 10
// 	d3 = n % 10
// 	hasil = (d1 < d2 and d2 < d3) or (d1 > d2 and d2 >  d3)
// 	output(hasil)
// endprogram
