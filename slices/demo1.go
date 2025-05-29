package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)

	fmt.Println(isimler)
	isimler[0] = "ali"
	isimler[1] = "veli"
	isimler[2] = "mehmet"
	isimler = append(isimler, "büşra")
	fmt.Println(isimler)

}
