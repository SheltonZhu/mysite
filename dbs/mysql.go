package dbs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var MysqlClient *gorm.DB

func init() {
	var err error
	MysqlClient, err = gorm.Open(mysql.Open(os.Getenv("MYSQL")), &gorm.Config{})

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if MysqlClient.Error != nil {
		fmt.Printf("database error %v", MysqlClient.Error)
	}
}
