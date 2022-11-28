package main

import (
	"fmt"
	"sync"
)
var wg = sync.WaitGroup{}
func main() {
	wg.Add(2)
	ch := make(chan string)
	go sortAlpha(ch)
	go printAlpha(ch)
	wg.Wait()
}
func sortAlpha(ch chan string) {
	for l := byte('a'); l < byte('l'); l++ {
		ch <- string(l)
	}
	wg.Done()
}
func printAlpha(ch chan string) {
	for l := byte('a'); l < byte('l'); l++ {
	fmt.Println(<-ch)
	}
	wg.Done()
}