package main

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)
func main(){
	db, err := sql.Open("mysql","root:11111111@/testdb")
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
}