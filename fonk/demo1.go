package fonk

import "fmt"

func Topla(sayi1 int, sayi2 int) {
	var toplam = sayi1 + sayi2
	fmt.Println("sonuç: ", toplam)
}
func Carp(sayi1 int, sayi2 int) int {
	var carp = sayi1 * sayi2
	return carp
}
