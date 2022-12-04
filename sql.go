package main
import (
	"database/sql"
	"fmt"
	"log"
	
	_ "github.com/go-sql-driver/mysql"
)
//var Num int = 123
func sql0(){
	db, err := sql.Open("mysql","root:11111111@/testdb")
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	/*
	creat table is not exists 'users'(
		'id' int(100) not null auto_increment,
		'username' varchar(45) not null,
		'password' varchar(40) not null,
		'isActive' tinyint(1) default null,
		primary key ('id'),
		unique key 'id_unique' ('id')
	)
	*/
	query := `creat table is not exists 'users'(
		'id' int(100) not null auto_increment,
		'username' varchar(45) not null,
		'password' varchar(40) not null,
		'isActive' tinyint(1) default null,
		primary key ('id'),
		unique key 'id_unique' ('id')
	)`
	_, err = db.Exec(query)
	if err != nil{
		log.Fatal(err.Error())
	}
	fmt.Println("blah blah ")
}