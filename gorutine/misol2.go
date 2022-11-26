package main
 
import (
    "fmt"
    "time"
)
 
func main() {
    fmt.Println("Before sleep the time is:", time.Now().Unix())     // Before sleep the time is: 1257894000
    time.Sleep(2 * time.Second)                                     // pauses execution for 2 seconds
    fmt.Println("After sleep the time is:", time.Now().Unix())      // After sleep the time is: 1257894002
}