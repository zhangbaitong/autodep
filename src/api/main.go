package main

import (
	"api/handlers"
	"api/rtop"
	"encoding/json"
	"fmt"
	"github.com/googollee/go-socket.io"
	"net/http"
	"runtime"
	"strings"
	"time"
)

type Monotor struct {
	Machine_ip string
	Ssh_port   int
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	go socketServer()
	httpServer()
}

//启动http服务
func httpServer() {
	fmt.Println("Server is start at ", time.Now().String(), " , on port 8090")
	http.HandleFunc("/useage", handlers.Useage)
	http.HandleFunc("/v1/", handlers.API_V1)
	http.ListenAndServe(":8090", nil)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("Server start faild error:", err)
	}
}

//启动socket服务
func socketServer() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		fmt.Println(err)
	}
	server.On("connection", func(so socketio.Socket) {
		fmt.Println("on connection")
		so.On("monitoring", func(msg string) {
			fmt.Println("emit:", msg+"123")
			go monitor(so, msg)
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		fmt.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("../../web")))
	fmt.Println("Serving at localhost:5000...")
	fmt.Println(http.ListenAndServe(":5000", nil))
}

func monitor(so socketio.Socket, msg string) {
	var monitor Monotor
	err_json := json.Unmarshal([]byte(msg), &monitor)
	if err_json != nil {
		fmt.Println("json data decode faild :", err_json)
	}

	fmt.Println(monitor.Machine_ip, monitor.Ssh_port)
	flag := 0

	so.On("disconnection", func() {
		fmt.Println("on disconnect")
		flag = 1
	})

	so.On("monitorOther", func(msg string) {
		err_json := json.Unmarshal([]byte(msg), &monitor)
		if err_json != nil {
			fmt.Println("json data decode faild :", err_json)
		}
	})

	sendMonitorInfo(so, monitor.Machine_ip, monitor.Ssh_port)

	timer := time.Tick(1 * time.Second)

	for {
		if flag == 1 {
			fmt.Println("on disconnect2")
			return
		}

		select {
		case <-timer:
			sendMonitorInfo(so, monitor.Machine_ip, monitor.Ssh_port)
		}
	}

}

func sendMonitorInfo(so socketio.Socket, machine_ip string, ssh_port int) {
	_, stats := rtop.GetHostStats(machine_ip, ssh_port)
	rtop.DisplayStats(stats)
	rets := strings.Replace(rtop.StringStats(stats), " ", "&nbsp;", -1)
	rets = strings.Replace(rets, "\n", "<br>", -1)
	fmt.Println("emit:", so.Emit("news", rets))
}
