package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Print("handler")
	if request.Method == "GET" {
		//返回上传页面
		data, err := ioutil.ReadFile("./static/view/upload.html")
		if err != nil {
			io.WriteString(writer, "interr success")
			return
		}
		io.WriteString(writer, string(data))
		return

	} else if request.Method == "POST" {
		//接收文件流，存储到本地目录
		file, head, err := request.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data err: %s\n", err.Error())
			return
		}
		defer file.Close()

		// "./static/tmp/" 这个目录之前就要存在
		newFile, err := os.Create("./static/tmp/" + head.Filename)

		if err != nil {
			fmt.Printf("Failed to ceate file  err: %s\n", err.Error())
			return
		}
		defer newFile.Close()
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to copy file  err: %s\n", err.Error())
			return
		}

		http.Redirect(writer, request, "/file/upload/suc", http.StatusFound)

	}
}
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finshed")
}
