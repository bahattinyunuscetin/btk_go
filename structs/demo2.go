package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("eklendi  ", c.firstName)
}
func (c customer) update() {
	fmt.Println("güncellendi  ", c.firstName)
}
func Demo2() {
	c := customer{firstName: "engin", lastName: "demiroğ", age: 35}
	c.save()
	c.update()
}
