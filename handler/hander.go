package handler

import (
	"io"
	"io/ioutil"
	"net/http"
)

func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		//返回上传页面
		data, err := ioutil.ReadFile("./static/view/upload.html")
		if err != nil {
			io.WriteString(writer, "interr success")
			return
		}

	} else if request.Method == "POST" {
		//接收文件流

	}
}
