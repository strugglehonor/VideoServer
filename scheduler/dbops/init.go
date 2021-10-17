package dbops

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/video_server/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&&parseTime=True",
		conf.DBUser, conf.DBPassWd, conf.DBHost, conf.Port, conf.DBName, conf.Charset)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// create table
	//db.Migrator().CreateTable(&Video_Delete_Record{})
}
