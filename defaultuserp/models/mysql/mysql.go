package mysql

import (
	"github.com/provider-go/go-logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func NewMysql(dsn string) {
	var err error
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("MySQL connection failed.", "err", err.Error())
	} else {
		logger.Info("MySQL connection successful!")
	}
}
