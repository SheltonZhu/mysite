package dao

import (
	"mysite/dbs"
)

type Bible struct {
	Id   int    `gorm:"primaryKey;autoIncrement;not null;column:bible_id" json:"id"`
	Text string `gorm:"unique;not null;column:bible_text" json:"text"`
}

func init() {
	dbs.MysqlClient.Table("bible_bible").AutoMigrate(&Bible{})
}

func GetBibleById(id int) (bible Bible, err error) {
	result := dbs.MysqlClient.Table("bible_bible").First(&bible, id)
	err = result.Error
	return
}

func ListBibles() (bibles []Bible, err error) {
	result := dbs.MysqlClient.Table("bible_bible").Find(&bibles)
	err = result.Error
	return
}

func (b *Bible) CreateBible() (int, error) {
	return CreateBible(b)
}

func CreateBible(bible *Bible) (int, error) {
	result := dbs.MysqlClient.Table("bible_bible").Create(bible)
	return bible.Id, result.Error
}

func DeleteBible(id int) (int, error) {
	result := dbs.MysqlClient.Table("bible_bible").Delete(&Bible{}, id)
	return int(result.RowsAffected), result.Error
}
