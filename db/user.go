package db

import (
	mydb "fileserver/db/mysql"
	"fmt"
)

// 用户注册
func UserSignup(username, passwd string) bool {
	stmt, err := mydb.DBConn().Prepare("insert ignore into tbl_user(`user_name`,`user_pwd`)values(?,?)")
	if err != nil {
		fmt.Println("Failed to insert tbl_user in prepare,err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, passwd)
	if err != nil {
		fmt.Println("Failed to insert tbl_user in exec ,err:" + err.Error())
		return false
	}

	if rowsAffected, err := ret.RowsAffected(); err == nil && rowsAffected > 0 {
		return true
	}
	return false
}

// 用户登录
func UserSignin(username, enc_passwd string) bool {
	stmt, err := mydb.DBConn().Prepare("select user_pwd from tbl_user where user_name = ? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else if rows == nil {
		fmt.Println("username not found:" + username)
		return false
	}
	var pwd string
	if rows.Next() {
		err = rows.Scan(&pwd)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Password for user %s is %s\n", username, pwd)
	} else {
		fmt.Println("User not found")
	}
	// 获取的rows 游标位于第一条数据之前，所以获取其中数据需要先 调用 next函数使游标位于第一条数据上
	// var pwd string
	// if err = rows.Scan(&pwd); err != nil {
	// 	fmt.Println(err.Error())
	// 	return false
	// }

	if pwd == enc_passwd {
		return true
	}
	return false

}

// 更新用户token
func UpdateToken(username, token string) bool {
	stmt, err := mydb.DBConn().Prepare("replace into tbl_user_token(`user_name`,`user_token`)values(?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
