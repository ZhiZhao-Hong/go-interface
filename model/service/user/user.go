package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BiSvc struct {
	*gin.Context
}

func NewBiSvc(c *gin.Context) *BiSvc {
	return &BiSvc{c}
}

type (
	GetUserInfoReq struct {
	}
	GetUserInfoResp struct {
	}
)

func (bv *BiSvc) GetUserInfo(tx *gorm.DB, req *GetUserInfoReq) (resp GetUserInfoResp, code int32, err error) {
	return
}
