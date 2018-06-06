package models

import (
	"os"

	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var (
	db   *gorm.DB
	err  error
	once *sync.Once
)

func Init() {
	db, err = gorm.Open("postgres", "host=127.0.0.1 port=9092 user=admin dbname=fushui password=123456")
	if err != nil {
		log.Errorf("open database failed : ", err.Error())
		os.Exit(-1)
	}
	return
}

func closedb() {
	db.Close()
}

func Close() {
	once.Do(closedb)
}
