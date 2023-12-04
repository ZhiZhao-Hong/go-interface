package base

import (
	"github.com/gin-gonic/gin"
	"go-interface/config"
	"go-interface/model/db"
	"go-interface/pkg/consts"
	"go-interface/pkg/jwt"
	"gorm.io/gorm"
)

type Option func(*Handler) error

type Handler struct {
	c   *config.Config
	db  *gorm.DB
	jwt *jwt.JwtMid
}

func New(c *config.Config, opts ...Option) (*Handler, error) {
	h := &Handler{
		c: c,
	}
	for _, opt := range opts {
		if err := opt(h); err != nil {
			return nil, err
		}
	}
	return h, nil
}

func InitMysql(c *db.Config) Option {
	return func(handler *Handler) error {
		conn, err := db.New(c)
		if err == nil {
			handler.db = conn
			handler.db = handler.db.Debug()
		}
		return err
	}
}

func InitJwt() Option {
	return func(handler *Handler) error {
		handler.jwt = jwt.NewJwtMid()
		return nil
	}
}

func (h *Handler) GetDB() *gorm.DB {
	return h.db
}

func (h *Handler) GetJwt() *jwt.JwtMid {
	return h.jwt
}

func (h *Handler) GetConfig() *config.Config {
	return h.c
}

func (h *Handler) GinH(c int32, err error, args ...interface{}) gin.H {
	msg := consts.GetMessage(c)
	var data interface{}
	var total interface{}
	if len(args) > 0 {
		data = args[0]
	}

	if len(args) > 1 {
		total = args[1]
	}

	return gin.H{
		"version": "",
		"code":    c,
		"msg":     msg,
		"total":   total,
		"data":    data,
	}
}
