package handler

import (
	"go-interface/config"
	"go-interface/handler/user"
	"go-interface/pkg/base"
)

type Handler struct {
	*base.Handler
	C *config.Config

	UserHandler *user.Handler
}

func New(c *config.Config, opts ...base.Option) (*Handler, error) {
	bh, err := base.New(c, opts...)
	if err != nil {
		return nil, err
	}
	h := &Handler{
		Handler:     bh,
		C:           c,
		UserHandler: user.New(bh),
	}
	return h, nil
}
