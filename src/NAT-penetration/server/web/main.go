package main

import (
	"encoding/json"
	"go_project/src/NAT-penetration/define"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		q := request.URL.Query()
		marshal, err := json.Marshal(q)
		if err != nil {
			panic(err)
		}
		_, err = writer.Write(marshal)
		if err != nil {
			panic(err)
		}
	})
	err := http.ListenAndServe(define.LocalHostAddr, nil)
	if err != nil {
		panic(err)
	}
}
