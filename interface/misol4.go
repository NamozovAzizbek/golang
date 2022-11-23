package main 
import "fmt"

// interface
type Shape interface {
  area() float32
}

 // struct to implement interface
type Rectangle struct {
  length, breadth float32
}

// use struct to implement area() of interface
func (r Rectangle) area() float32 {
  return r.length * r.breadth
}

// access method of the interface
func calculate(s Shape) {
  fmt.Println("Area:", s.area())
}

// main function
func main() {
 
  // assigns value to struct members
  rect := Rectangle{7, 4}

  // call calculate() with struct variable rect
  calculate(rect)
    
}