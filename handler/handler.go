package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
*
处理文件上传
*/
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传HTML页面
		data, err := os.ReadFile("./static/view/upload.html")
		if err != nil {
			io.WriteString(w, "Internel server Error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接受文件流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Fail to get data,err:%s", err.Error())
			return
		}
		defer file.Close()

		//创建本地文件接受文件流
		newFile, err := os.Create("C:/Users/Administrator/GolandProjects/file-server/tmp/" + head.Filename)
		if err != nil {
			fmt.Println("Fail to create file,err:%s", err.Error())
			return
		}
		defer newFile.Close()
		//3、拷贝文件
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Println("Fail to copy file,err:%s", err.Error())
		}
		//4、上传成功则重定向
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

/*
*
上传已完成
*/
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload Success!")
}
