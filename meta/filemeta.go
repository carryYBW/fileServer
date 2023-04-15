package meta

// 文件元信息结构
type FileMeta struct {
	//文件的唯一标识
	FileSha1 string
	Filename string
	FileSize string
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
}

// 通过 fileSha1 获取文件的元信息
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}