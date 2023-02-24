package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

// 支持配置key默认值设置
// 支持读取JSON,TOML,YAML,HCL,envfile和java properties等多种不同类型配置文件
// 可以监听配置文件的变化，并重新加载配置文件
// 读取系统环境变量的值
// 读取存储在远程配置中心的配置数据，如ectd，Consul,firestore等系统，并监听配置的变化
// 从命令行读取配置
// 从buffer读取配置
// 可以显示设置配置的值

func main() {
	viper.SetConfigFile("conf.yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	addr := viper.Get("addr")
	fmt.Println(addr)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().Format("2006-01-02 15:04:05")
		w.Write([]byte(t))
	})
	err := http.ListenAndServe(addr.(string), nil)
	if err != nil {
		fmt.Println(err)
	}

}
