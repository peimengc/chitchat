package config

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
	"os"
	"sync"
)

type App struct {
	Address  string
	Static   string
	Log      string
	Locale   string
	Language string
}

type Database struct {
	Driver   string
	Address  string
	Database string
	User     string
	Password string
}

type Configuration struct {
	App          App
	Db           Database
	LocaleBundle *i18n.Bundle
}

var config *Configuration
var once sync.Once

func LoadConfig() *Configuration {
	//单例模式
	once.Do(func() {
		//打开文件
		file, err := os.Open("config.json")
		if err != nil {
			log.Fatalln("无法读取配置文件", err)
		}
		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err != nil {
			log.Fatalln("配置加载失败", err)
		}
		//本地化
		bundle := i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
		bundle.MustLoadMessageFile(config.App.Locale + "/active.en.json")
		bundle.MustLoadMessageFile(config.App.Locale + "/active." + config.App.Language + ".json")
		config.LocaleBundle = bundle
	})
	return config
}
