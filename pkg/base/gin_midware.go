package base

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MidwareDemo(tx *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}
