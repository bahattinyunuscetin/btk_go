package errorhandling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 50
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arası gir")
	}
	if tahmin > aklimdakiSayi {
		return "daha kücük gir", nil
	}
	if tahmin < aklimdakiSayi {
		return "daha büyük gir", nil
	}
	if tahmin == aklimdakiSayi {
		return "bildin tebrikler", nil
	}
	return "bildin", nil
}
func Demo2() {
	mesaj, hata := TahminEt(101)
	fmt.Println(mesaj)
	fmt.Println(hata)
}
