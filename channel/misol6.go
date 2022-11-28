package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
func main() {
	 in := make( chan int, 1)
	in <- 4
	//wg.Add(1)
	go isFunc(in)
	fmt.Println("sakjfkajsfkj", <-in)
	wg.Wait()
}
func isFunc(in chan int){
	num := <- in
	num = num+1
	fmt.Println(num)
	defer wg.Done()
}