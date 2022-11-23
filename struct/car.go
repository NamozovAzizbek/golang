package main

import "fmt"

type Car struct {
	name  string
	price int
	year  int
}

func main() {
	c1 := Car{name: "BMW", price: 45000, year: 2020}
	fmt.Println(c1)
}