package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}



type Labels struct {
	HostName string `gorm:"unique;not null"`
	Label string
}

type Actions struct {
	Time time.Time
	Context string
}

type Logs struct {
	Types string
	Time time.Time
	Context string
	Node string
}




func CloseDB(db *gorm.DB) {
	err := db.Close()
	if err != nil{
		panic(err)
	}
}

func Setup() {
	var err error
	db, err = gorm.Open("sqlite3", "../db.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Model{},&Nodes{},&Tokens{},&Labels{},&Actions{},&Logs{})

}