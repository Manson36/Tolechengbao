package datasource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type pqDB struct {
	*gorm.DB
}

//使用gorm框架初始化 postgres连接
func (p *pqDB)initDB() {
	arg := "host=127.0.0.1 port=5432 user=postgres password=123456 dbname=postgres sslmode=disable"
	db, err := gorm.Open("postgres", arg)
	if err != nil {
		msg := fmt.Sprintf("init postgres db error: %s", err.Error())
		panic(msg)
	}

	if err = db.DB().Ping(); err != nil {
		msg := fmt.Sprintf("ping postgres db error: %s", err.Error())
		panic(msg)
	}

	db.LogMode(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(time.Hour)
	db.DB().SetMaxOpenConns(30)

	p.DB = db
}

var PqDB = &pqDB{}