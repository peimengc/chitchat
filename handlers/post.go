package handlers

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/peimengc/chitchat/models"
	"net/http"
)

func PostThread(w http.ResponseWriter, r *http.Request) {
	//登录验证
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		//根据session获取用户
		user, err := session.User()
		if err != nil {
			danger("未找到用户", err)
		}
		//解析Form参数
		err = r.ParseForm()
		if err != nil {
			danger("参数异常", err)
		}
		//获取话题
		thread, err := models.ThreadByUuid(r.PostFormValue("uuid"))

		if err != nil {
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "thread_not_found",
			})
			errorMessage(w, r, msg)
		} else {
			//创建
			if _, err := user.CreatePost(thread, r.PostFormValue("body")); err != nil {
				errorMessage(w, r, "回复失败")
			}

			http.Redirect(w, r, fmt.Sprint("/thread/read/?id=", thread.Uuid), http.StatusFound)
		}

	}

}
