package goroutines

import (
	"fmt"
	"time"
)

func Ciftsayilar() {
	for i := 0; i < 10; i += 2 {
		fmt.Println("çift sayi:", i)
		time.Sleep(1 * time.Second)
	}
}

func Teksayilar() {
	for i := 1; i < 10; i += 2 {
		fmt.Println("tek sayi:", i)
		time.Sleep(1 * time.Second)
	}
}
