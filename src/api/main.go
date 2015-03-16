package main

import (
	"api/handlers"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println("Server is start at ", time.Now().String(), " , on port 8080")
	http.HandleFunc("/useage", handlers.Useage)
	http.HandleFunc("/v1/", handlers.API_V1)
	http.ListenAndServe(":8080", nil)
}
