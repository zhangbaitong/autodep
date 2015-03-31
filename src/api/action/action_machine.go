package action

import (
	"api/common"
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

const (
	dbName = "../db/autodep.db"
	//dbName = "./autodep.db"
)

func RegisterMachine(request common.RequestData) string {
	var params Machine
	jsonErr := json.Unmarshal([]byte(request.Params), &params)
	if jsonErr != nil {
		logger.Println("json data decode faild :", jsonErr)
	}

	if "" == strings.Trim(params.Machine_name, " ") {
		return "Machine_name can't be empty"
	}

	if "" == strings.Trim(params.Machine_ip, " ") {
		return "Machine_ip can't be empty"
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into machines(machine_name,machine_ip,docker_port,is_use,remark) values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(params.Machine_name, params.Machine_ip, params.Docker_port, params.Is_use, params.Remark)

	if err != nil {
		log.Fatal("参数1", err)
	}

	tx.Commit()

	return "ok"
}

type Machine struct {
	Machine_name string
	Machine_ip   string
	Docker_port  int
	Is_use       int
	Remark       string
}
