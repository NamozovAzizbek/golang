package main

import "fmt"

type Animal struct {
	name   string
	type_a string
	age    int
}

func PraAnimal(name, type_a string, age int) *Animal {
	n := new(Animal)
	n.name = name
	n.type_a = type_a
	n.age = age
	return n
}
func main() {
	a1 := PraAnimal("lion", "wild", 11)
	fmt.Println(a1)
}