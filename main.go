package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("文件系统")

	http.HandleFunc("file/upload", handler.uploadHandler)
}
