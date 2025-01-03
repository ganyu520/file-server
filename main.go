package main

import (
	"fmt"
	"net/http"

	"github.com/ganyu520/file-server/handler"
)

func main() {
	//路由规则
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	//端口监听
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
		return
	}
}
