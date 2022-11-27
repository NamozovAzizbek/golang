package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {
	go numbers()
	go alphabets()
	fmt.Println("Before sleep the time is:", time.Now().Unix())     // Before sleep the time is: 1257894000
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("After sleep the time is:", time.Now().Unix())      // After sleep the time is: 1257894002
	fmt.Println("main terminated")
}