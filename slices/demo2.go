package slices

import "fmt"

func Demo2() {
	sehirler := []string{"ankara", "istanbul", "izmir"}
	fmt.Println(sehirler)
	sehirlerkopya := make([]string, len(sehirler))
	fmt.Println(sehirlerkopya)
	copy(sehirlerkopya, sehirler)
	fmt.Println(sehirlerkopya)

}
