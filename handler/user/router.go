package user

import "github.com/gin-gonic/gin"

func (h *Handler) InitUserRouter(api *gin.RouterGroup) {
	wx := api.Group("/user")
	wx.Use(h.GetJwt().Auth())
	{
		wx.GET("", h.GetUserInfo)
	}
}
