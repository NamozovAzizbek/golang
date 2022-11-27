package main

import "fmt"

func main() {
	m := make(chan string)
	n := make(chan int)
	go funcv(n, m)
	fmt.Println(<-n)
	fmt.Println(<-m)
}
func funcv(n chan int, m chan string) {
	n <- 13
	m <- "send massage"
	fmt.Println("func ended")
}