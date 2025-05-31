package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string) // 1.tanımlama yöntemi
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"
	fmt.Println(sozluk["book"])
	fmt.Println("eleman sayisi : ", len(sozluk))
	fmt.Println("sözlük:", sozluk)
	delete(sozluk, "book")
	fmt.Println("eleman sayısı :", len(sozluk))
	fmt.Println("sözlük:", sozluk)
	deger, varMi := sozluk["table"]
	fmt.Println(deger)
	fmt.Println("listede olma durumu: ", varMi)
	k2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"} //2. tanımlanma yöntemş
	fmt.Println(k2)
}
