package main

import "fmt"

func main() {
	ch := make(chan int)
	//ch <- 9
	go printNumber(ch)
	x :=<-ch
	x++
	fmt.Println("func main", x)
}
func printNumber(ch chan int){
	ch <- 7
	fmt.Println("printNumber")
}
