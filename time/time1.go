package main

import (
	"fmt"
	"time"
)

func main() {
	(time.Date(2020, time.October, 12, 12, 2, 21, 2, time.UTC))
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Local().Year())
	//fmt.Println(time.Before(u))
}
