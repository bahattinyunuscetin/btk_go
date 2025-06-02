package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt")
	//nill
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("dosya bulunamadı", pErr.Path)
			return
		} else {
			fmt.Println("dosya bulunamadı")
			return
		}

	} else {
		fmt.Println(f.Name())
	}
}
