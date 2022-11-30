package main
import ("fmt"
"time"
)

func main(){
start := time.Now()
a := 5
b := a + 2
fmt.Println(b, time.Since(start),time.Now().Unix(), time.Now())
}