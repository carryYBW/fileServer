package db

import (
	"database/sql"
	mydb "fileserver/db/mysql"
	"fmt"
)

// 文件上传完成，保存meta到数据库
func OnFileUploadFinished(filehash string, filename string,
	filesize int64, fileaddr string) bool {
	// 生成预备指令
	stmt, err := mydb.DBConn().Prepare(
		// "insert ignore into tbl_file(`file_sha1`,`file_name`,`file_size`,`file_addr`,`status` values(?,?,?,?,1))"
		// insert ignore 是 插入数据的key重复则忽略
		"INSERT IGNORE INTO tbl_file (file_sha1, file_name, file_size, file_addr, status) VALUES (?, ?, ?, ?, 1)")
	if err != nil {
		fmt.Println("failed to prepare statement ,err:" + err.Error())
		return false
	}

	defer stmt.Close()
	// 执行
	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	// 判断执行结果
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("File with hash:%s has been upload before", filehash)
		}
		return true
	}
	return false
}

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

// 获取文件元信息
func GetFileMeta(fileHash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1,file_name,file_size,file_addr from tbl_file where file_sha1=? and status=1 limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()
	tfile := TableFile{}
	err = stmt.QueryRow(fileHash).Scan(&tfile.FileHash, &tfile.FileName, &tfile.FileSize, &tfile.FileAddr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &tfile, nil
}
