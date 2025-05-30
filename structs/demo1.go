package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"bilgisayar", 500, "XYZ"})

}

type product struct {
	name      string
	unitprice float64
	brand     string
}
