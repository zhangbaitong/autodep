package main

import (
	"database/sql"
	"fmt"
	_ "github.com/wendal/go-oci8"
	"log"
)

type Users struct {
	UserId int
	Uname  string
}

func main() {
	log.Println("Oracle Driver Connecting....")
	//用户名/密码@实例名 如system/123456@orcl、sys/123456@orcl
	db, err := sql.Open("oci8", "BOOKMAN/password@orcl")
	if err != nil {
		log.Fatal(err)
		panic("数据库连接失败")
	} else {
		defer db.Close()
		var users []Users = make([]Users, 0)
		rows, err := db.Query("select * from users")
		if err != nil {
			log.Fatal(err)
		} else {
			for rows.Next() {
				var u Users
				rows.Scan(&u.UserId, &u.Uname)
				users = append(users, u)
			}
			fmt.Println(users)
			defer rows.Close()
		}

	}

}
