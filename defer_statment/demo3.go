package defer_statment

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) Save() {
	fmt.Println("ürün kaydedildi :", p.productName)
	Log()
}
func Log() {
	fmt.Println("loglandı")
}
func Demo3() {
	p := product{productName: "laptop", unitPrice: 5000}
	defer p.Save()
	p = product{productName: "mouse", unitPrice: 45}
	fmt.Println("işlem başarılı")
}
