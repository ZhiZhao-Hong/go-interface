package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go-interface/model/db"
	"go-interface/pkg/log"
)

type Option func(config *Config)

type Service struct {
	HttpPort         uint  `json:"httpPort"`
	TokenFailureTime int64 `json:"tokenFailureTime"`
}

type Config struct {
	DB      *db.Config `json:"DB"`
	Log     *log.Conf  `json:"log"`
	Service *Service   `json:"service"`
}

func NewConfig(fps []string, options ...Option) (*Config, error) {
	if len(fps) < 1 {
		return nil, fmt.Errorf("配置文件不能为空")
	}
	for _, fp := range fps {
		viper.SetConfigFile(fp)
		if err := viper.MergeInConfig(); err != nil {
			return nil, err
		}
	}
	c := new(Config)
	if err := viper.Unmarshal(c); err != nil {
		return nil, err
	}
	for _, option := range options {
		option(c)
	}
	return c, nil
}

func SetLogger() Option {
	return func(config *Config) {
		config.Log.InitLogger()
	}
}
