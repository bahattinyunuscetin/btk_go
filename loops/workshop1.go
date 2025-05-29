package loops

import "fmt"

func Demo3() {
	aklimdakisayi := 80
	puan := 100
	var tahminedilensayi int

	fmt.Println("1 ile 100 arasında bir sayı tahmin edin")

	// İlk tahmin için geçerli giriş yapılana kadar döngü
	for {
		fmt.Print("Tahmininiz: ")
		fmt.Scanln(&tahminedilensayi)
		if tahminedilensayi >= 1 && tahminedilensayi <= 100 {
			break
		}
		fmt.Println("Lütfen 1 ile 100 arasında bir sayı girin!")
	}

	// Tahmin kontrol döngüsü
	for aklimdakisayi != tahminedilensayi && puan > 0 {
		if tahminedilensayi < aklimdakisayi {
			fmt.Println("Daha büyük bir sayı girin.")
		} else {
			fmt.Println("Daha küçük bir sayı girin.")
		}
		puan -= 10

		// Yeni tahmin için geçerli giriş kontrolü
		for {
			fmt.Print("Yeni tahmininiz: ")
			fmt.Scanln(&tahminedilensayi)
			if tahminedilensayi >= 1 && tahminedilensayi <= 100 {
				break
			}
			fmt.Println("Lütfen 1 ile 100 arasında bir sayı girin!")
		}
	}

	if aklimdakisayi == tahminedilensayi {
		fmt.Println("Bravo, kazandınız!")
		fmt.Printf("Puanınız: %d\n", puan)
	} else {
		fmt.Println("Kaybettiniz. Puanınız 0.")
		fmt.Printf("%v", puan)
	}
}
