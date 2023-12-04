package service

import (
	"fmt"
	"go-interface/config"
	"go-interface/handler"
	"go-interface/pkg/base"
	"go-interface/pkg/jwt"
	"go-interface/pkg/log"
)

type ServiceContainer struct {
	Config *config.Config
}

func (sc *ServiceContainer) Init() {
	sc.initJwt()
	sc.initGin()
}

func (sc *ServiceContainer) initJwt() {
	jwt.Init(sc.Config.Service.TokenFailureTime)
}

func (sc *ServiceContainer) initGin() {
	var h *handler.Handler
	var err error
	// 注册web插件
	h, err = handler.New(
		sc.Config,
		base.InitMysql(sc.Config.DB),
		base.InitJwt())
	if err != nil {
		log.Fatal(err.Error())
	}
	// 注册路由
	r := h.InitRouter()
	// 启动服务
	if err = r.Run(fmt.Sprintf(":%d", sc.Config.Service.HttpPort)); err != nil {
		log.Fatal(err.Error())
	}
}
