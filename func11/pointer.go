package main

import "fmt"

func main() {
	var a, b int
	a = 2
	b = 3
	swap(&a, &b)
	fmt.Println(a, b)
}
func swap(a, b *int) {
	var c int
	c = *a
	*a = *b
	*b = c
}