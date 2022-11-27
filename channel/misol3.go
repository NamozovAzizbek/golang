package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 9
	fmt.Println(<-ch)
	go printNumber(ch)
	fmt.Println("func main")
}
func printNumber(ch chan int){
	ch <- 7
	fmt.Println("printNumber")
}
