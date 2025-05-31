package main

import (
	"fmt"
	"golesson/channels"
)

func main() {

	CiftSayiCN := make(chan int)
	teksayicn := make(chan int)

	go channels.Teksayilar(teksayicn)
	go channels.Ciftsayilar(CiftSayiCN)

	ciftsayitoplam, teksayitoplam := <-CiftSayiCN, <-teksayicn
	carpim := ciftsayitoplam * teksayitoplam
	fmt.Println("main bitti", carpim)
}
