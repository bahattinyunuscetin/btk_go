package variables

import "fmt"

func Demo3() {
	var sayi1, sayi2, sayi3 int = 10, 23, 45
	var enbuyuk int = sayi1

	if sayi2 > enbuyuk {
		enbuyuk = sayi2
	}
	if sayi3 > enbuyuk {
		enbuyuk = sayi3
	}

	fmt.Printf("en büyük sayı : %v", enbuyuk)

}
