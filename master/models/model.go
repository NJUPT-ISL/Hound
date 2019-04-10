package models

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

type Tokens struct {
	HostName string `gorm:"unique;not null"`
	Token string
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

func OpenDB() *gorm.DB{
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func AutoMigrate(db *gorm.DB){
	db.AutoMigrate(&Nodes{},&Tokens{},&Labels{},&Actions{},&Logs{})
}


func CloseDB(db *gorm.DB) {
	err := db.Close()
	if err != nil{
		panic(err)
	}
}
