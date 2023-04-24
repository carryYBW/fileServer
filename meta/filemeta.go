package meta

import (
	mydb "fileserver/db"
	"fmt"
)

// 文件元信息结构
type FileMeta struct {
	//文件的唯一标识
	FileSha1 string
	Filename string
	FileSize int64
	Location string
	UploadAt string
}

// 存储所有的 元文件信息   key 是 FileSha1  文件的唯一标识
var fileMetas map[string]FileMeta

// init() 会在 程序运行初始化的时候执行一次
func init() {
	fileMetas = make(map[string]FileMeta)
}

// 新增/更新文件元信息
func UpdateFileMeta(f FileMeta) {
	fileMetas[f.FileSha1] = f
	fmt.Println(f)
}

// 新增/更新文件元信息到 数据库
func UpdateFileMetaDB(f FileMeta) bool {
	return mydb.OnFileUploadFinished(f.FileSha1, f.Filename, f.FileSize, f.Location)
}

// 通过 fileSha1 获取文件的元信息
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// 通过 fileSha1 获取文件的元信息
func GetFileMetaDB(fileSha1 string) (FileMeta, error) {
	tfile, err := mydb.GetFileMeta(fileSha1)

	if err != nil {
		return FileMeta{}, err
	}
	fMeta := FileMeta{}
	fMeta.FileSha1 = tfile.FileHash
	fMeta.Filename = tfile.FileName.String
	fMeta.FileSize = tfile.FileSize.Int64
	fMeta.Location = tfile.FileAddr.String
	return fMeta, nil
}

// 删除  实际环境中 要考虑 多线程
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
