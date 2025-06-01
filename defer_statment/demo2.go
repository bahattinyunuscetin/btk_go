package defer_statment

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		return "çift sayidir"
	}

	if sayi%2 != 0 {
		return "tek sayidir"
	}

	return "belli değil "
}

func Test() {
	sonuc := Demo2(9)
	fmt.Println(sonuc)
}
func DeferFunc() {
	fmt.Println("bitti")

}
