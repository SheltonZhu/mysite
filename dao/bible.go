package dao

import (
	"fmt"
	"github.com/SheltonZhu/mysite/dbs"
	"log"
)

const BibleTableName = "bible_bible"

type Bible struct {
	Id   int    `gorm:"primaryKey;autoIncrement;not null;column:bible_id" json:"id"`
	Text string `gorm:"unique;not null;column:bible_text" json:"text"`
}

func init() {
	if err := dbs.MysqlClient.Table(BibleTableName).AutoMigrate(&Bible{}); err != nil {
		log.Println(fmt.Sprintf("migrate faild. %v", err))
	}
}

func GetBibleById(id int) (bible Bible, err error) {
	result := dbs.MysqlClient.Table(BibleTableName).First(&bible, id)
	err = result.Error
	return
}

func ListBibles() (bibles []Bible, err error) {
	result := dbs.MysqlClient.Table(BibleTableName).Find(&bibles)
	err = result.Error
	return
}

func (b *Bible) CreateBible() (int, error) {
	return CreateBible(b)
}

func CreateBible(bible *Bible) (int, error) {
	result := dbs.MysqlClient.Table(BibleTableName).Create(bible)
	return bible.Id, result.Error
}

func DeleteBible(id int) (int, error) {
	result := dbs.MysqlClient.Table(BibleTableName).Delete(&Bible{}, id)
	return int(result.RowsAffected), result.Error
}
