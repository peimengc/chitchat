package handlers

import (
	"errors"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	. "github.com/peimengc/chitchat/config"
	"github.com/peimengc/chitchat/models"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	logger    *log.Logger
	config    *Configuration
	localizer *i18n.Localizer
)

func init() {
	config = LoadConfig()
	localizer = i18n.NewLocalizer(config.LocaleBundle, config.App.Language)

	file, err := os.OpenFile("logs/chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("无法创建日志文件", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

//通过cookie判断用户是否登录
func session(w http.ResponseWriter, r *http.Request) (session models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		session = models.Session{Uuid: cookie.Value}
		if ok, _ := session.Check(); !ok {
			err = errors.New("session 失效")
		}
	}
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s/%s.html", config.App.Language, file))
	}
	templates := template.Must(template.ParseFiles(files...))

	templates.ExecuteTemplate(w, "layout", data)
}

// 返回版本号
func Version() string {
	return "0.1"
}

func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

func errorMessage(w http.ResponseWriter, r *http.Request, msg string) {
	http.Redirect(w, r, fmt.Sprint("/err?msg=", msg), http.StatusFound)
}
