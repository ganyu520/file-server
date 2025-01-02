package main

import (
	"file-server/handler"
	"fmt"
	"net/http"
)

func main() {
	//路由规则
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	//端口监听
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server, err:%s", err.Error())
		return
	}
}
