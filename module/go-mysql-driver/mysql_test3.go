package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	outStmt,err := db.Prepare("SELECT * FROM squarenum WHERE number > ?")
	if err != nil{
		fmt.Println(err.Error())
	}
	rows, err :=outStmt.Query(0)
	if err != nil{
		fmt.Println(err.Error())
	}

	var col1,col2 int
	for rows.Next(){
		err := rows.Scan(&col1,&col2)
		if err != nil{
			fmt.Println(err.Error())
			continue
		}
		fmt.Println(col1, col2)
	}
}