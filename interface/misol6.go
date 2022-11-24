package main

import "fmt"

type Shape interface {
	area() float32
	permanent() float32
}

type Rectangle struct {
	lenght, breadth float32
}

// permanent implements Shape
func (Rectangle) permanent() float32 {
	panic("unimplemented")
}

func (r Rectangle) area() float32 {
	return r.lenght * r.breadth
}

type Triangle struct {
	base, height float32
}

func (t Triangle) area() float32 {
	return 0.5 * t.base * t.height
}
func calculate(s Shape) {
	fmt.Println(s.area())
}
func main() {
	rec := Rectangle{5, 4}
	//fmt.Println(rec.area())
	t := Triangle{5, 4}
	// shaps := []Shape{rec, t}
	// for _, sh := range shaps {
	// 	calculate(sh)
	// }
	calculate(rec)
	fmt.Println(t.area())
}
