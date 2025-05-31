package channels

import (
	"fmt"
	"time"
)

func Ciftsayilar(CiftSayiCN chan int) {
	toplam := 0

	for i := 0; i < 10; i += 2 {

		toplam = toplam + i

		fmt.Println("çift sayı toplama çalışıyor")
		time.Sleep(1 * time.Second)
	}
	CiftSayiCN <- toplam
}

func Teksayilar(teksayicn chan int) {
	toplam := 0
	for i := 1; i < 10; i += 2 {

		toplam = toplam + i
		fmt.Println("tek sayı toplama çalışıyor")
		time.Sleep(1 * time.Second)

	}
	teksayicn <- toplam
}
