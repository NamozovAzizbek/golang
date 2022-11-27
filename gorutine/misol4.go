// Program to run two goroutines concurrently

package main
import (
  "fmt"
  "time"
)

// create a function
func display(message string) {

  // infinite for loop
  for {
    fmt.Println(message)

    // to sleep the process for 1 second
    time.Sleep(time.Second * 1)
  }
}

func main() {

  // call function with goroutine
  go display("Process 1")

  display("Process 2")

}