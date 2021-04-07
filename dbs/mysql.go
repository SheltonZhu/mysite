package dbs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mysite/conf"
	"os"
)

var MysqlClient *gorm.DB

func init() {
	var err error
	var mysqlConn = ""
	if os.Getenv("MYSQL") != "" {
		mysqlConn = os.Getenv("MYSQL")
	} else {
		mysqlConn = conf.Conf.Db.Mysql.ConnStr
	}
	MysqlClient, err = gorm.Open(mysql.Open(mysqlConn), &gorm.Config{})

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if MysqlClient.Error != nil {
		fmt.Printf("database error %v", MysqlClient.Error)
	}
}
