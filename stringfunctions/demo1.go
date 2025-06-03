package stringfunctions

import (
	"fmt"

	s "strings"
)

// case sensitive  küçük büyük harf duyarlı
func Demo1() {
	isim := "ali"
	fmt.Println(s.Count(isim, "g"))
	fmt.Println(s.Contains(isim, "n")) // mevcutmu true fulse
	fmt.Println(s.Index(isim, "n"))    // ilk gördüğ n nin index ini dödürür   bulamazsa -1 döndürür

}
