package main

import "fmt"

func main() {
	a := []int{1, 3, 4, 5, 6}
	fmt.Println(topsum(a...))
}
func topsum(nums ...int ) int {
	sum := 0
	for _, v := range nums{
		sum += v
	}
	return sum
}