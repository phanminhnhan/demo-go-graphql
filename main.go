package main

import (
	"fmt"
	"graphql_postrgres/database"
)

func main(){
	sql := &database.Sql{
		Host:     "localhost",
		Port:      5432,
		Username: "postgres",
		Password: "postgres",
		DbName:   "graphql",
	}
	sql.ConnectDB()
	defer sql.Close()
	//var user  = model.User{}
	//var demo = Demo{}
	result, err := sql.Db.Query("select * from demo where name = $1", "nhan2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(*result)

}
type Demo struct {
	name string
	age int

}