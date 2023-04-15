package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type Sha1Stream struct {
	_sha1 hash.Hash
}

func (obj *Sha1Stream) Updata(data []byte) {
	if obj._sha1 == nil {
		obj._sha1 = sha1.New()
	}
	obj._sha1.Write(data)
}

//这是一个使用 Go 语言编写的函数，它属于 Sha1Stream 结构体的方法。
//该方法使用 SHA1 哈希算法对输入数据进行哈希计算，并返回计算结果的十六进制字符串表示。

//具体来说，
//该方法首先调用 _sha1.Sum([]byte(""))，将一个空的字节数组作为参数传递给 _sha1.Sum() 方法，
//计算出输入数据的 SHA1 哈希值。
//然后，它使用 Go 语言的 hex 包中的 EncodeToString() 函数将哈希值转换为十六进制字符串，并将其作为函数返回值返回。

// 需要注意的是，在函数返回的字符串中，每个字节都被转换为了两个十六进制数字，因此字符串的长度是原始哈希值长度的两倍。
func (obj *Sha1Stream) Sum() string {
	return hex.EncodeToString(obj._sha1.Sum([]byte("")))
}

func Sha1(data []byte) string {
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

func FileSha1(file *os.File) string {
	_sha1 := sha1.New()
	io.Copy(_sha1, file)
	return hex.EncodeToString(_sha1.Sum(nil))
}

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func FileMD5(file *os.File) string {
	_md5 := md5.New()
	io.Copy(_md5, file)
	return hex.EncodeToString(_md5.Sum(nil))
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err

}

func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, info fs.FileInfo, err error) error {
		result = info.Size()
		return nil
	})
	return result
}
