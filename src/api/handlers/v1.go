package handlers

import (
	"api/common"
	"api/action"
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("access http://localhost:8080/useage for how to use this api.")
}

func API_V1(w http.ResponseWriter, r *http.Request) {
	strMethod := r.URL.Path[len("/v1/"):]
	if err=r.ParseForm(); err != nil {
                    fmt.Println("Server start faild error:", err)
	}

                strPostData:=r.FormValue("request")
                //fmt.Println(strPostData)
                //var request common.RequestData                
                var req interface{} 
                err := json.Unmarshal([]byte(strPostData), &req)               
	if  err != nil {
                    fmt.Println("json data decode faild :", err)
	}

                request, _ := req.(map[string]interface{})
                common.DisplayJson(request)
	var ret string
	switch strMethod {
                    case "fig/create":
                                {
                                    ret = action.FigCreate(request)
                                }
                    case "echo":
                                {
                                    ret = action.Actionecho()
                                }
                    case "version":
                                {
                                    ret = action.Actionversion()
                                }
	}
	v1 := common.Response{Method: strMethod,Code:0,Messgae:"ok", Data: ret}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, err1 := json.Marshal(v1)
	if err1 != nil {
		fmt.Println(err1)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(body)
}
