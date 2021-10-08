package dbops

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/video_server/conf"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	dsn := fmt.Sprintf("gorm:gorm@tcp(%s:%d)/gorm?charset=%s&&parseTime=True", conf.DBHost, conf.Port, conf.Charset)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}