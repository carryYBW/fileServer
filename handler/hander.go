package handler

import (
	"io/ioutil"
	"net/http"
)

func loadfile(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		ioutil.ReadFile("./static/view/index.html")
	} else if request.Method == "POST" {

	}
}
