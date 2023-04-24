package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	//匿名导入的方式  导入后 会进行一个初始化，并将自己注册到 database/sql中
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:YBW@1good@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Fialed to connect to mysql,err:" + err.Error())
		os.Exit(1)
	}
}

// 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}
