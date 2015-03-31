package handlers

import (
	"api/action"
	"api/common"
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("access http://localhost:8080/useage for how to use this api.")
}

func API_V1(w http.ResponseWriter, r *http.Request) {

	strMethod := r.URL.Path[len("/v1/"):]
	if err = r.ParseForm(); err != nil {
		common.Log().Println("Server internal error:", err)
	}

	//var request map[string]interface{}
	var request common.RequestData

	if r.Method == "GET" {
	} else {
		strPostData := r.FormValue("request")
		//var request common.RequestData

		err := json.Unmarshal([]byte(strPostData), &request)
		if err != nil {
			fmt.Println("json data decode faild :", err)
		}
		//request, _ = req.(map[string]interface{})
		//common.DisplayJson(request)
	}

	var ret string
	switch strMethod {
	case "fig/create":
		{
			ret = action.FigCreate(request)
		}
	case "container/create":
		{
			ret = action.CreateContainer(request)
		}
	case "container/inspect":
		{
			ret = action.InspectContainer(request)
		}
	case "container/list":
		{
			ret = action.ListContainers(request)
		}
	case "container/changes":
		{
			ret = action.ContainerChanges(request)
		}
	case "container/stop":
		{
			ret = action.StopContainer(request)
		}
	case "container/restart":
		{
			ret = action.RestartContainer(request)
		}
	case "container/pause":
		{
			ret = action.PauseContainer(request)
		}
	case "container/unpause":
		{
			ret = action.UnpauseContainer(request)
		}
	case "container/kill":
		{
			ret = action.KillContainer(request)
		}
	case "container/info":
		{
			// ret = action.InfoContainer(request)
		}
	case "version":
		{
			//ret = action.Actionversion()
		}
	case "image/list":
		{
			ret = action.ListImages(request)
		}
	case "reg/list":
		{
			// ret = action.ActionRegList()
		}
	case "reg/tags":
		{
			// ret = action.ActionRegTags(r.Form["ns"], r.Form["rep"])
		}
	case "reg/search":
		{
			// ret = action.ActionRegSearch(r.Form["q"], r.Form["n"], r.Form["page"])
		}
	case "reg/info":
		{
			 ret = action.ActionAllInfo(request)
		}
	case "machine/register":
		{
			ret = action.RegisterMachine(request)
		}
	case "machine/search":
		{
			ret = action.SearchMachine(request)
		}
	}
	v1 := common.Response{Method: strMethod, Code: 0, Messgae: "ok", Data: ret}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	body, err1 := json.Marshal(v1)
	if err1 != nil {
		fmt.Println(err1)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(body)
}
