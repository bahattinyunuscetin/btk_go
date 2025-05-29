package variables

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("para yetersiz")
	} else {
		fmt.Println("paranız hazırlanıyor")
	}

}
