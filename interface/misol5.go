package main 
import "fmt"

// interface
type Shape interface {
  area() float32
  perimeter() float32
}

 // Rectangle struct implements the interface
type Rectangle struct {
  length, breadth float32
}

// Rectangle provides implementation for area()
func (r Rectangle) area() float32 {
  return r.length * r.breadth
}

// access method of the interface
func calculate(s Shape) float32 {
  return s.area()
}

// main function
func main() {
 
  // assigns value to struct members
  r := Rectangle{7, 4}

  // call calculate() with struct variable rect 
  rectangleArea := calculate(r)
  fmt.Println("Area of Rectangle:", rectangleArea)
}