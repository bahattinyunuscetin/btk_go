package main

import (
	"fmt"
	"golesson/refactoring"
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
	//fmt.Println(errorhandling.TahminEt2(102))
	//restful.Demo2()
	// project_.AddProduct()
	// project_.GetAllProducts()
	product, _ := refactoring.AddProduct()
	fmt.Println(product)

	products, _ := refactoring.GetAllProducts()
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)

	}

}
