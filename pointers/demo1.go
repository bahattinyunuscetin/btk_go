package pointers

import "fmt"

func Demo1(sayi *int) {
	*sayi = *sayi + 1
	fmt.Println("demodaki sayi", sayi)

}
func Test() {
	sayi := 10
	Demo1(sayi)

}
