package main

import (
	"fmt"
	"net/http"

	//导入包对应的文件路径,fileserver 为 go mod init moulde名，在go.mod文件中有显示
	handler "fileserver/handler1"
)

func main() {
	fmt.Print("文件系统")

	//第一个路径需要添加 `/`
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	http.HandleFunc("/user/signup", handler.SignupHandler)
	http.HandleFunc("/user/signin", handler.SigninHandler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Printf("failed to start server ,err: %s", err.Error())
	}
}
