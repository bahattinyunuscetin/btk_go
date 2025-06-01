package interface_

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	addres             string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}
type CreditCalculater interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 12
}
func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 12
}
func CalculateMonthlyPayment(credits []CreditCalculater) float64 {
	var paymentTotal float64 = 0
	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal
}
func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, addres: "ankara"}
	credit2 := Mortgage{rate: 10, creditPaymentTotal: 20000, addres: "istanbul"}
	credit3 := Car{rate: 15, creditPaymentTotal: 60000, carInfo: "polo"}

	credits := []CreditCalculater{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)

	fmt.Println(total)
}
