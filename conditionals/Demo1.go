package conditionals

import "fmt"

func main() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("para yetersiz")
	}
	if cekilmekIstenen <= hesap {
		fmt.Println("paraniz hazirlaniyor")
		hesap = hesap - cekilmekIstenen
	}
	fmt.Println("Bitti. hesaptaki para : " + fmt.Sprintf("%v", hesap))

}
