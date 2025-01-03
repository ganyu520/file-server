package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ganyu520/file-server/meta"
	"github.com/ganyu520/file-server/util"
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
		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "/root/autodl-tmp/file-server/tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		//创建本地文件接受文件流
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Fail to create file,err:%s", err.Error())
			return
		}
		defer newFile.Close()
		//3、拷贝文件
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Fail to copy file,err:%s", err.Error())
			return
		}

		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)
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

// GetFileMetaHandler:获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
