#  BTK Akademi - Go (Golang) Dili Kursu Projeleri

Merhaba yazılımcı dostum! Bu repo, [BTK Akademi](https://www.btkakademi.gov.tr/portal/course/go-ile-programlamaya-giris-12760) uzerinden tamamladigim **Go ile Programlamaya Giriş** kursu kapsaminda gelistirdigim proje ve ornekleri barindiriyor. "Go neymis ya" diyorsan, iceri gir bir bak: sade, hizli, guclu.

Go diline yeni baslayanlar, "Ben ne yazsam?" diyenler ya da "Sadece run edip calissin yeter" kafasinda olanlar icin birebir.

---

## 🔍 Go (Golang) Nedir?

Go (aka **Golang**), 2007'de Google tarafindan yazilim dunyasinin en kronik sikayetlerini cozmek icin dogdu. 2009'da open-source yapildi, 2012'de 1.0 ile olgunluga erdi. "Statik tipli ama okuması kolay bir dil yazalim, C gibi hizli olsun ama JavaScript kadar basit dursun" kafasiyla gelistirildi.

### Neden bu kadar seviliyor? Çünkü:

* 🚀 **Işık hızında derleme & calistirma**
* 🔄 **Goroutine + Channel** ile native concurrency destegi
* 🧼 Sade, temiz ve okunabilir syntax (tipki IKEA mobilyası gibi: az parça, bol iş)
* 🛠 Gömülü toolchain: `go fmt`, `go test`, `go doc`, her şey kutudan çıkar çıkmaz hazır
* ⚙️ Docker, Kubernetes gibi dev sistemlerin dili — altyapının temel taşı!

---

## 📁 Proje Klasörleri

```bash
.
├── 01-Temel-Kavramlar/         # Değişkenler, veri tipleri, operatörler
├── 02-Kosullar-Donguler/       # if-else, switch-case, for döngüleri
├── 03-Fonksiyonlar/            # Parametreli ve döndüren fonksiyonlar
├── 04-Veri-Yapilari/           # Dizi, slice, map kullanımı
├── 05-Struct-Interface/        # Struct ve interface tanımları
├── 06-GoRoutines-Channels/     # Eşzamanlılık örnekleri
├── 07-Web-Uygulamalari/        # Basit HTTP sunucusu ve API örnekleri
└── README.md
```

Her klasor, konuyla ilgili minik ama ogretici ornekler icerir. Yaz, calistir, oyna, boz, tekrar yaz. Yazilim boyle ogrenilir!

---

## 🧠 Öne Çıkan Konular

### ✅ Veri Yapıları

```go
type Person struct {
    Name string
    Age  int
}

var ages = map[string]int{"Ali": 25, "Veli": 30}
```

### ✅ Fonksiyonlar ve Arayüzler (Interfaces)

```go
func Topla(a int, b int) int {
    return a + b
}

type Hayvan interface {
    SesCikar() string
}
```

### ✅ Goroutine ve Channel Kullanımı

```go
go yazdir("Merhaba") // eşzamanlı çalışır

ch := make(chan string)
ch <- "veri"
msg := <-ch
```

### ✅ Basit Web Sunucusu

```go
package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Merhaba Go dünyası!")
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
```

Tarayıcıya `http://localhost:8080` yaz, Go seni selamlasın 👋

---

## 🛠 Kurulum ve Kullanım

1. [Go resmi sitesi](https://go.dev/dl/) üzerinden sistemine uygun Go sürümünü indir.
2. Terminalde proje klasorune gir.
3. Calistirmak icin:

```bash
cd 01-Temel-Kavramlar/
go run main.go
```

Hepsi bu. Derlenir, calisir, patlamaz. Go'nun mottosu: "simple, reliable, efficient" — cok iddiali ama dogru.

---

## 🌍 Go ile Geliştirilen Dev Projeler

| Proje          | Açıklama                                       |
| -------------- | ---------------------------------------------- |
| **Docker**     | Container teknolojisinin öncüsü                |
| **Kubernetes** | Modern uygulama orkestrasyonu                  |
| **Terraform**  | Altyapı yönetimi (infrastructure as code)      |
| **Hugo**       | Static site generator (blog yazarları bayılır) |
| **Caddy**      | Otomatik SSL destekli web sunucusu             |

Go'nun gücü sadece yazılımda değil; internetin yapı taşlarını taşır hale geldi.

---

## 💡 Neden Go Öğrenmelisin?

* 👶 Yeni başlayanlar için basit, kıdemliler için sağlam
* 🧵 Native concurrency ile çoklu işlemde uçuyor
* 👑 Google destekli: uzun ömürlü ve güvenli bir yatırım
* 📦 Tek binary ile dağıtım: dependency bağımlılığı derdi yok
* 🧘‍♂️ "Keep it simple" felsefesini iliklerine kadar yaşatıyor

> "Go dili sade, hızlı ve eşzamanlı programlamaya uygun yapısıyla geleceğin sistem programlama dillerinden biridir."

---

## 🤝 Katkıda Bulunmak İstersen

Bu repo eğitim amaçlıdır. Sen de projelere katkı sağlamak, farklı çözümler eklemek veya yorum yapmak istersen forkla, kodunu yaz, PR gönder. Açık kaynak ruhu yaşasın! 🙌

---

## ✍️ Hazırlayan

**Bahattin Yunus Çetin**
🎓 KTÜ Yazılım Mühendisliği Öğrencisi
🚀 Go & Python meraklısı


---

> Kodun temiz, compiler hatasız, channel'ların tıkanmasın!

README'yi okuduysan, Go diline adım atmaya çoktan hazırsın demektir. Hadi bakalım, `go run` ile macera başlasın! 🧑‍💻
