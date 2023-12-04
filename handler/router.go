package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-interface/pkg/log"
)

func (h *Handler) InitRouter() *gin.Engine {
	r := gin.Default()
	cfg := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowCredentials: false,
		AllowHeaders:     []string{"authorization", "content-type", "*"},
	}

	// 控件
	r.Use(cors.New(cfg))
	r.Use(log.GinToLog())

	// 路由器注册
	api := r.Group("/api")
	h.UserHandler.InitUserRouter(api)
	return r
}
