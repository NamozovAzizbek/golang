// Program to illustrate multiple goroutines

package main

import (
  "fmt"
  "time"
)

func display(message string) {

  fmt.Println(message)

}

func main() {
  fmt.Println("Salom Gorutine")

  // run two different goroutine
  go display("Process 1")
  go display("Process 2")
  time.Sleep(time.Second * 1)
  go display("Process 3")

  // to sleep main goroutine for 1 sec
}