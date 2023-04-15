package handler

import (
	"encoding/json"
	"fileserver/meta"
	"fileserver/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// 文件上传
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

		//初始化 filemeta信息
		fileMeta := meta.FileMeta{
			Filename: head.Filename,
			Location: "./static/tmp/" + head.Filename,
			// UploadAt: time.Now().Format("2000-01-01 24:00:00"),这个有问题
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
			//
		}
		//os.Getwd() 可以获取当前工作目录，是多变的，问价的存储应该选择 绝对路径  实际项目中可以使用  os.mkdir 创建一个绝对路径文件目录。在存储文件
		//这里先使用  选相对路径，存储在项目路径下 便于查看

		// "./static/tmp/" 这个目录之前就要存在
		// newFile, err := os.Create("./static/tmp/" + head.Filename)
		newFile, err := os.Create(fileMeta.Location)

		if err != nil {
			fmt.Printf("Failed to ceate file  err: %s\n", err.Error())
			return
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to copy file  err: %s\n", err.Error())
			return
		}
		//将文件指针 复原  指向文件开头
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		//完成filemeta 信息的填充
		meta.UpdateFileMeta(fileMeta)

		http.Redirect(writer, request, "/file/upload/suc", http.StatusFound)

	}
}

// 上传已完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finshed")
}

// 获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filehash := r.Form["filehash"][0]
	fmeta := meta.GetFileMeta(filehash)
	data, err := json.Marshal(fmeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// 文件下载接口
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fm := meta.GetFileMeta(fsha1)

	f, err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	data, err := ioutil.ReadFile(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//浏览器识别是文件，可以下载。。这通常用于传输文件或其他二进制数据。当客户端收到这个响应时，它会将响应体中的数据保存为一个二进制文件，而不会尝试解析它。
	w.Header().Set("Content-Type", "application/octect-stream")
	//设置文件名
	w.Header().Set("Content-Disposition", "attachment; filename="+fm.Filename)
	w.Write(data)

}

// 文件修改  重命名
func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	opType := r.Form.Get("op")
	fileSha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	curFileMeta := meta.GetFileMeta(fileSha1)
	curFileMeta.Filename = newFileName
	meta.UpdateFileMeta(curFileMeta)

	data, err := json.Marshal(curFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

// 文件删除
func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileSha1 := r.Form.Get("filehash")

	fmeta := meta.GetFileMeta(fileSha1)

	// 物理删除
	os.Remove(fmeta.Location)
	//逻辑删除
	meta.RemoveFileMeta(fileSha1)

	w.WriteHeader(http.StatusOK)
}
