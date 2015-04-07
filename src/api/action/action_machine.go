package action

import (
	"api/common"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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
		logger.Println(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		logger.Println(err)
	}
	stmt, err := tx.Prepare("insert into machines(machine_name,machine_ip,docker_port,is_use,remark) values(?, ?, ?, ?, ?)")
	if err != nil {
		logger.Println(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(params.Machine_name, params.Machine_ip, params.Docker_port, params.Is_use, params.Remark)

	if err != nil {
		logger.Println("参数1", err)
	}

	tx.Commit()

	return "ok"
}

func SearchMachine(request common.RequestData) string {
	var params Machine
	jsonErr := json.Unmarshal([]byte(request.Params), &params)
	if jsonErr != nil {
		logger.Println("json data decode faild :", jsonErr)
		return "faild"
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		logger.Println(err)
		return "faild"
	}
	defer db.Close()

	where := fmt.Sprintf("where is_use=%d ", params.Is_use)

	if "" != strings.Trim(params.Machine_name, " ") {
		where += fmt.Sprintf(" and machine_name like '%%%s%%' ", params.Machine_name)
	}

	if "" != strings.Trim(params.Machine_ip, " ") {
		where += fmt.Sprintf(" and machine_ip like '%%%s%%' ", params.Machine_ip)
	}

	if params.Docker_port > 0 {
		where += fmt.Sprintf(" and docker_port = %d ", params.Docker_port)
	}

	rows, err := db.Query("select machine_name, machine_ip, docker_port, is_use, remark from machines " + where)
	if err != nil {
		logger.Println(err)
		return "faild"
	}
	defer rows.Close()
	var machines []Machine = make([]Machine, 0)
	for rows.Next() {
		var m Machine
		rows.Scan(&m.Machine_name, &m.Machine_ip, &m.Docker_port, &m.Is_use, &m.Remark)
		machines = append(machines, m)
	}

	strMachines, err := json.Marshal(machines)
	if err != nil {
		logger.Println(err)
		return "faild"
	}

	return string(strMachines)
}

type Machine struct {
	Machine_name string
	Machine_ip   string
	Docker_port  int
	Is_use       int
	Remark       string
}

type MachineSearch struct {
	Machine_name string
	Machine_ip   string
	Docker_port  int
	Is_use       int
}
