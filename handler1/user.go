package handler

import (
	dblayer "fileserver/db"
	"fileserver/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	pwd_salt = "#901"
)

// 处理用户注册请求
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")

	//简单的校验
	if len(username) < 3 || len(passwd) < 5 {
		w.Write([]byte("Invalid parameter,len is not enough"))
		return
	}
	enc_passwd := util.Sha1([]byte(passwd + pwd_salt))
	suc := dblayer.UserSignup(username, enc_passwd) // layer 层
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}

// 用户登录
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	enc_passwd := util.Sha1([]byte(password + pwd_salt))
	//1.校验用户名及密码
	pwdChecked := dblayer.UserSignin(username, enc_passwd)
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}

	//2.token 验证 或 session和cookie
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username, token)
	if !upRes {
		w.Write([]byte("FAILED"))
		return
	}
	//3.登录成功后，重定向到首页

	// w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "http://" + r.Host + "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	w.Write(resp.JSONBytes())

}

// 查询用户信息
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	//1.解析参数
	r.ParseForm()
	username := r.Form.Get("username")
	token := r.Form.Get("token")
	//2.校验token
	if !isTokenValid(token) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	//3.查询用户信息
	user, err := dblayer.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	resp := util.RespMsg{
		Code: 0,
		Msg:  "ok",
		Data: user,
	}
	w.Write(resp.JSONBytes())

}

// 生产token
func GenToken(username string) string {
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_toknesalt"))
	return tokenPrefix + ts[:8]
}

func isTokenValid(token string) bool {
	if len(token) != 40 {
		return false
	}
	//1.判断token时效性，是否过期，这里是 token的后8位是个时间
	//2.从数据库中查询
	//3.对比是否一致
	return true
}
