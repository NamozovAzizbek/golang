package main

import "fmt"

type Laptop struct {
	name  string
	price int
	ssd   string
}

func main() {
	var l Laptop
	l.name = "HP"
	l.price = 410
	l.ssd = "M2 256 GB"
	fmt.Println(l)
}
