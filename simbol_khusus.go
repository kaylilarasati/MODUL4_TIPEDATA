package main

import "fmt"

func main() {
	var a, b rune
	var hasilA, hasilB bool

	fmt.Scanf("%c%c", &a, &b)

	hasilA = !((a >= 'A' && a <= 'Z') ||
		(a >= 'a' && a <= 'z') ||
		(a >= '0' && a <= '9'))

	hasilB = !((b >= 'A' && b <= 'Z') ||
		(b >= 'a' && b <= 'z') ||
		(b >= '0' && b <= '9'))

	fmt.Println(hasilA, hasilB)
}
