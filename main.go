package main

import (
	"fmt"
	errorhandling "golesson/error_handling"
)

func main() {

	// CiftSayiCN := make(chan int)
	// teksayicn := make(chan int)

	// go channels.Teksayilar(teksayicn)
	// go channels.Ciftsayilar(CiftSayiCN)

	// ciftsayitoplam, teksayitoplam := <-CiftSayiCN, <-teksayicn
	// carpim := ciftsayitoplam * teksayitoplam
	// fmt.Println("main bitti", carpim)
	// interface_.Demo2()
	fmt.Println(errorhandling.TahminEt2(102))

}
