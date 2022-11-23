package main

import "fmt"

type User struct {
	name, surname string
	ID, age       int
	Pay           Payment
}
type Payment struct {
	salery, bonus int
}
func  (u User) GetFulname() string {
	return u.name + " " + u.surname
}
func (u Payment) GetMoney() int{
	return u.salery + u.bonus
}
func main() {
	u1 := User{
		name:    "Azizbek",
		surname: "Namozov",
		ID:      1,
		age:     22,
		Pay:     Payment{salery: 1000, bonus: 300},
	}
	fmt.Println(u1.GetFulname())
	fmt.Println(u1.Pay.GetMoney())
}
