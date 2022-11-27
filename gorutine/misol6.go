package main

import (
	"fmt"
	"sync"
)

func runner1(wg *sync.WaitGroup) {
	fmt.Println("\n I am first function")
	wg.Done()
}
func runner2(wg *sync.WaitGroup){
	fmt.Println("\n I am second function")
	wg.Done()
}
func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go runner1(wg)
	go runner2(wg)
	wg.Wait()
}

func main() {
	execute() 
	//time.Sleep(100*time.Millisecond)
}