package main
import "fmt"

func main()	{
	var jam, menit, detik, totalDetik int
	fmt.Print("Masukkan jumlah detik: ")
	fmt.Scanf("%d", &totalDetik)

	jam = totalDetik / 3600
	sisa := totalDetik % 3600
	menit = sisa / 60
	detik = sisa % 60

	fmt.Printf("Jam: %d\n", jam)
	fmt.Printf("Menit: %d\n", menit)
	fmt.Printf("Detik: %d\n", detik)
}
// kamus konversiWaktu
// Kamus : jam, menit, detik, totalDetik : int
// Algoritma
// 	Input (totalDetik)
//	jam = totalDetik / 3600
//	sisa = totalDetik % 3600
//	menit = sisa / 60
//	detik = sisa % 60
//	Output (jam, menit, detik)
// endprogram
