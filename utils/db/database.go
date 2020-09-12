package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gostart/utils/conf"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

func DBConn() *gorm.DB {
	var (
		dialect string
		driver = conf.GetDBDriver()
		dbHost = conf.GetDBHost(driver)
		dbPort = conf.GetDBPort(driver)
		dbName = conf.GetDBName(driver)
		dbUser = conf.GetDBUser(driver)
		dbPass = conf.GetDBPass(driver)
	)

	switch driver {
	case "mysql":
		dialect = fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	case "postgres":
		dialect = fmt.Sprintf("host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	default:
		dialect = fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	}

	db, err := gorm.Open(driver, dialect)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)

	return db
}