package main

import "fmt"

type greater interface {
	great(string) string
}

type uzb struct{}
type eng struct{}

func (u *uzb) great(name string) string {
	return fmt.Sprintf("Salom, salom %s", name)
}
func (e *eng) great(name string) string {
	return fmt.Sprintf("Hello, hi %s", name)
}

func sayHello(g greater, name string) {
	fmt.Print(g.great(name))

}

func main() {
	sayHello(&uzb{}, "azizbek")
	sayHello(&eng{}, "azizbek")

}
