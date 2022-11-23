package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"salom", "dunyo", "va ", "men"}
	ajrat(a...)

}
func ajrat(s ...string) {
	fmt.Println(strings.Join(s, "-"))
}