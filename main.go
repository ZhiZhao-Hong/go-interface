package main

import (
	"go-interface/config"
	"go-interface/service"
	"log"
	"os"
	"os/signal"
)

// 初始化环境变量
var (
	conf = func() *config.Config {
		c, err := config.NewConfig([]string{
			"./tomls/config.toml",
			"./tomls/secret.toml",
		}, config.SetLogger())
		if err != nil {
			log.Fatal(err.Error())
		}
		return c
	}()
)

func Run() {
	// 启动服务
	sc := service.ServiceContainer{Config: conf}
	sc.Init()
	// 退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutdown server ...")
	log.Println("server exited")
	os.Exit(0)
}

func main() {
	Run()
}
