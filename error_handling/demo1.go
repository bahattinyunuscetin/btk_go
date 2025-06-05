package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt") //dosyayı açmaya çalışır  açamazsa nil=err olur
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok { //Eğer hata bir *os.PathError türündeyse, onu pErr olarak al.
			fmt.Println("dosya bulunamadı", pErr.Path) //Eğer gerçekten PathError ise, yani dosya yolu ile ilgili bir hata varsa, hata mesajı ile birlikte hangi dosyada hata olduğunu yazdır
			return
		} else {
			fmt.Println("dosya bulunamadı")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}
