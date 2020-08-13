package main

import (
	. "github.com/peimengc/chitchat/config"
	. "github.com/peimengc/chitchat/routes"
	"log"
	"net/http"
)

func main() {
	startWebServer("8080")
}

func startWebServer(port string) {
	//加载配置文件
	config := LoadConfig()

	r := NewRouter()

	//静态资源文件
	assets := http.FileServer(http.Dir(config.App.Static))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)

	log.Printf("在%s开启http服务\r\n", config.App.Address)

	err := http.ListenAndServe(config.App.Address, nil)

	if err != nil {
		log.Printf("在%s开启http服务失败", config.App.Address)
		log.Println("error:", err)
	}

}
