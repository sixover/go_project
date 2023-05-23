package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type addParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type resData struct {
	Data int `json:"data"`
	Code int `json:"code"`
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	//解析数据
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("it is something wrong"))
	}
	var param addParam
	json.Unmarshal(req, &param)

	//处理业务
	ret := param.Y + param.X
	//返回响应
	resp, _ := json.Marshal(resData{
		Data: ret,
		Code: 0,
	})
	w.Write(resp)
}

func main() {
	http.HandleFunc("/add", addHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
