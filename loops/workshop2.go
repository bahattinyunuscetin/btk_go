package loops

import "fmt"

func Demo4() {
	sayi := 0
	fmt.Println("bir sayı gir")
	fmt.Scanln(&sayi)
	asalMİ := true
	for i := 2; i < sayi/2; i++ {
		if sayi%i == 0 {
			asalMİ = false
		}

	}
	if asalMİ {
		fmt.Println("sayi asaldir")

	} else {
		fmt.Println("sayi asal degildir ")
	}

}
