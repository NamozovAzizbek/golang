package main

import "fmt"

func main() {
	ch := make(chan string)
	go reciveDate(ch)
	fmt.Println("No date. Recive blocked")
	ch <- "Date Recive. Operation Sacssefull"
}
func reciveDate(ch chan string) {
	fmt.Println(<-ch)
}