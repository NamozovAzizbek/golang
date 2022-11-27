package main

import "fmt"


func main() {
	ch := make(chan int, 4)
	 go printMy(ch)
	for i := 0; i < 4; i ++{

		fmt.Println(<-ch)
	}
}
func printMy(ch chan int){
	for i := 0; i < 4; i ++{
		ch <- i
	}
}