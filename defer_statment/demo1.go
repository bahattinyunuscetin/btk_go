package defer_statment

import "fmt"

func A() {
	fmt.Println("a çaıştı")
}
func C() {
	fmt.Println("c çaıştı")
}
func B() {
	defer A() // en son defer e giren ilk çıkar
	defer C()
	fmt.Println("b çalıştı")
	A()
	C()
}
