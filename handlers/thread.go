package handlers

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/peimengc/chitchat/models"
	"net/http"
)

func NewThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "auth.navbar", "new.thread")
	}
}

func CreateThread(w http.ResponseWriter, r *http.Request) {
	//验证是否登录
	session, err := session(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		//解析表单
		err = r.ParseForm()
		if err != nil {
			fmt.Println("无法解析表单")
		}
		//获取用户信息
		user, err := session.User()
		if err != nil {
			fmt.Println("用户不存在")
		}
		//创建
		if _, err = user.CreateThread(r.PostFormValue("topic")); err != nil {
			fmt.Println("话题创建失败:", err)
		}
		//返回主页
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func ReadThread(w http.ResponseWriter, r *http.Request) {
	//格式化url query
	queryVal := r.URL.Query()
	//查询thread
	thread, err := models.ThreadByUuid(queryVal.Get("id"))
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "thread_not_found",
		})
		errorMessage(w, r, msg)
	} else {
		//是否登录
		_, err = session(w, r)
		//加载模板
		if err != nil {
			generateHTML(w, &thread, "layout", "navbar", "thread")
		} else {
			generateHTML(w, &thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}
