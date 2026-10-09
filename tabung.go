package main

import "fmt"
type Tabung struct {
	jariJari, tinggi int
	luas, volume float64
	luasAlas, luasDinding float64
}	

func main() {
	const pi = 3.14
	var t Tabung
	fmt.Scan(&t.jariJari, &t.tinggi)
	t.luasAlas = pi * float64(t.jariJari) * float64(t.jariJari)
	t.luasDinding = 2 * pi * float64(t.jariJari) * float64(t.tinggi)
	t.luas = 2 * t.luasAlas + t.luasDinding	
	t.volume = t.luasAlas * float64(t.tinggi)
	fmt.Printf("%.2f\n", t.luas)
	fmt.Printf("%.2f\n", t.volume)
} 