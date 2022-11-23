package main

import (
	"fmt"
)

type Human struct {
	name  string
	sname string
	age   int
}
func NewHuman() Human{
	var n  Human
	return n
} 
func ParamtrHuman(name, sname string , age int) *Human{
	n := new(Human)
	n.age = age
	n.name = name
	n.sname = sname
	return n
}
func main() {
	 //a := Human {name : "Azizbek", sname: "Namozov", age: 22}
	 var b Human
	 b.age = 54
	 b.sname = "blah"
	 n := NewHuman()
	fmt.Println(n)
	h5 := ParamtrHuman("Asror", "Nayimov", 45)
	fmt.Println(h5)
}