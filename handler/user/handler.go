package user

import (
	"github.com/gin-gonic/gin"
	"go-interface/model/service/user"
	"go-interface/pkg/base"
	"go-interface/pkg/consts"
	"net/http"
)

type Handler struct {
	*base.Handler
}

func New(bh *base.Handler) *Handler {
	h := &Handler{
		bh,
	}
	return h
}

func (h *Handler) GetUserInfo(c *gin.Context) {
	var code int32
	var err error
	var data interface{}
	var count int64
	defer func() { c.JSON(http.StatusOK, h.GinH(code, err, data, count)) }()
	var req user.GetUserInfoReq
	if err = c.ShouldBind(&req); err != nil {
		code = consts.CODE_PARAM
		return
	}
	bv := user.NewBiSvc(c)
	data, code, err = bv.GetUserInfo(h.GetDB(), &req)
}
