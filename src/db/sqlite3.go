package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const (
	dbName = "./sqlite3.db"
)

type Users struct {
	UserId int
	Uname  string
}

func main() {
	os.Remove("./sqlite3.db")

	db, err := sql.Open("sqlite3", dbName)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := `create table users (userId integer, uname text);`
	db.Exec(sql)
	sql = `insert into users(userId,uname) values(1,'Mike');`
	db.Exec(sql)
	sql = `insert into users(userId,uname) values(2,'John');`
	db.Exec(sql)
	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []Users = make([]Users, 0)
	for rows.Next() {
		var u Users
		rows.Scan(&u.UserId, &u.Uname)
		users = append(users, u)
	}
	fmt.Println(users)

}
