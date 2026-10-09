package main
import "fmt"

func main(){
	var berat, tinggi, bmi float64
	fmt.Scan(&berat, &tinggi)
	bmi = berat / (tinggi * tinggi)
	fmt.Printf("%.2f\n", bmi)
}
// program bmi
// kamus : berat, tinggi, bmi : float
// algoritma : 
// 	input(berat, tinggi)
// 	bmi = berat / (tinggi * tinggi)
// 	output(bmi)
// endprogram