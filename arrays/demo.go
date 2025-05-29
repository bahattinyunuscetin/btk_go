package arrays

import "fmt"

func Demo2() {
	sayilar := [5]int{20, 30, 50, 10, 2}
	fmt.Println(sayilar)
	max := sayilar[0]
	for i := 0; i < len(sayilar); i++ {
		if max < sayilar[i] {
			max = sayilar[i]
		}
	}
	fmt.Println(max)
}
