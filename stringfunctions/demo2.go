package stringfunctions

import (
	"fmt"

	s "strings"
)

func Demo2() {
	isim := "engin"
	fmt.Println(s.HasPrefix(isim, "eng")) // isim eng ile başlıyormu
	fmt.Println(s.Contains(isim, "n"))
	fmt.Println(s.Index(isim, "n"))
	harfler := []string{"e", "n", "g", "i", "n"}
	sonuc := s.Join(harfler, "*") //harfler in arasına * ekle
	fmt.Println(sonuc)
	fmt.Println(s.Replace(sonuc, "*", "+", -1)) //yıldızları + ile değiştir
	fmt.Println(s.Split(sonuc, "*"))            // * lar ı referas alarak metn sıralar  [e n g i n]
	fmt.Println(s.Repeat(sonuc, 5))             // 5 kere tekrarla
}
