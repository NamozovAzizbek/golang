package main

import "fmt"

var isConnect = false

func main() {
db()
fmt.Println("Status db : ", isConnect)
}
func db(){
	defer disconnect()
	fmt.Println("Status db : ", isConnect )
	connect()
	fmt.Println("Status db : ", isConnect)
}
func connect() {
	isConnect = true
	fmt.Println("Database connect !!!")
}
func disconnect(){
	isConnect = false
	fmt.Println("Database disconnected !!!")
}