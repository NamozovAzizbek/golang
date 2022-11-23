package main
import "fmt"

func main(){
f1 := fibo()
for i := 0; i < 10; i ++{
	fmt.Println(f1(i)," ", f1(-i*2))
}
}
func fibo() func(int) int {
	sum := 0
	return func (x int ) int {
		sum += x
		return sum
	}
}