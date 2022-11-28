package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	ch1 := make(chan string)
	go func() {
		for {

			ch <- "tez"
			time.Sleep(time.Millisecond * 1000)
		}
	}()
	go func() {
		for {

			ch1 <- "sekin"
			time.Sleep(time.Second * 2)
		}
	}()
	for {
		select {
		case m := <-ch1:
			fmt.Println(m)
		case n := <-ch:	
			fmt.Println(n)
		}
		fmt.Println("************************")
	}
}
