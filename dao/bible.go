package dao

import "mysite/dbs"

type Bible struct {
	Id   int    `gorm:"primaryKey;autoIncrement;not null;column:bible_id" json:"id"`
	Text string `gorm:"unique;not null;column:bible_text" json:"text"`
}

func init() {
	dbs.MysqlClient.Table("bible_bible").AutoMigrate(&Bible{})
}

func GetBibleById(id int) (bible Bible) {
	dbs.MysqlClient.Table("bible_bible").First(&bible, id)
	return
}

func ListBibles() (bibles []Bible) {
	dbs.MysqlClient.Table("bible_bible").Find(&bibles)
	return
}

func (b *Bible) CreateBible() (int, error) {
	return CreateBible(b)
}

func CreateBible(bible *Bible) (int, error) {
	result := dbs.MysqlClient.Create(bible)
	return bible.Id, result.Error
}
