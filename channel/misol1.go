package main
import "fmt"

func main() {

  // create channel
  ch := make(chan string)

  // function call with goroutine
  go receiveData(ch)

  fmt.Println("No data. Receive Operation Blocked")
  ch <- "Data Received. Receive Operation Successful" 

  // send data to the channel 

}

func receiveData(ch chan string) {
    
  // receive data from the channel
  fmt.Println(<-ch)

}