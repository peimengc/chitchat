package handlers

import (
	"github.com/peimengc/chitchat/models"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "auth.layout", "navbar", "login")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "auth.layout", "navbar", "signup")
}

func SignupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		errorMessage(w, r, "无法解析表单信息")
	}

	user := models.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}

	if err = user.Create(); err != nil {
		danger("用户注册失败", err)
		errorMessage(w, r, "用户注册失败")
	}

	http.Redirect(w, r, "/login", 302)

}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	//解析表单信息
	err := r.ParseForm()

	if err != nil {
		errorMessage(w, r, "无法解析表单信息")
	}

	//根据邮箱查找用户
	user, err := models.UserByEmail(r.PostFormValue("email"))

	if err != nil {
		errorMessage(w, r, "账号或密码错误")
	}

	//验证密码
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		//创建session
		session, err := user.CreateSession()

		if err != nil {
			danger("无法创建session", err)
			errorMessage(w, r, "登录失败")
		}

		//写入cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		})

		//重定向到主页
		http.Redirect(w, r, "/", http.StatusFound)

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//获取cookie
	cookie, err := r.Cookie("_cookie")

	if err != http.ErrNoCookie {
		//删除session表数据
		session := models.Session{Uuid: cookie.Value}

		session.DeleteByUuid()
	}

	//重定向到主页
	http.Redirect(w, r, "/", http.StatusFound)
}
