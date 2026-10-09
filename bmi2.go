package main
import "fmt"

func main(){
	var berat, tinggi, bmi float64
	fmt.Scan(&bmi, &tinggi)
	berat = bmi * (tinggi * tinggi)
	fmt.Printf("%.0f\n", berat)
}
