package main

import "fmt"

func main() {
	var durum bool
	var metin1 string = "ali"
	var metin2 string = "veli"
	durum = metin1 == metin2
	fmt.Println(durum)
	fmt.Println(!durum)
}
