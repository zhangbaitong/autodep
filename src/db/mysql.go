package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	UserId int
	Uname  string
}

func main() {
	//db, err := sql.Open("mysql", "user:password@/dbname")
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		fmt.Println("连接数据库失败")
	}
	defer db.Close()
	var users []Users = make([]Users, 0)
	sqlStr := "select * from users"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; rows.Next(); i++ {
			var u Users
			rows.Scan(&u.UserId, &u.Uname)
			users = append(users, u)
		}
		fmt.Println(users)
	}
}
