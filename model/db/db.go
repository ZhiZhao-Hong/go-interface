package db

import (
	"fmt"
	"go-interface/pkg/log"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Config struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
	IP       string `json:"ip,omitempty"`
	Port     int    `json:"port,omitempty"`
	DB       string `json:"db,omitempty"`
	//Dsn string `json:"dsn"`
}

func New(c *Config) (*gorm.DB, error) {
	w := NewZapLogger(log.Get())
	newLogger := logger.New(w, logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      false,
	})
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		c.UserName, c.Password, c.IP, c.Port, c.DB)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
}

type ZapLogger struct {
	*zap.SugaredLogger
}

func (logger *ZapLogger) Printf(format string, args ...interface{}) {
	logger.SugaredLogger.Infof(format, args...)
}

func NewZapLogger(logger *zap.SugaredLogger) *ZapLogger {
	return &ZapLogger{
		logger,
	}
}
