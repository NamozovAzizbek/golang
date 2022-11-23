package main
import "fmt"

func main(){
a := []int {3, 4, 5, 7, 6}
void(a...)
}
func void( a ...int){
	for _, v := range a{
		fmt.Print(v)
	}
}