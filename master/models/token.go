package models

import (
	"time"
)

type Token struct {
	Name         string `gorm:"PRIMARY_KEY;unique;not null"`
	Token        string
	GenerateTime time.Time
	UpdateTime   time.Time
}

func CheckToken(name string) (err error, ok bool) {
	t, err := GetToken(name)
	if err != nil {
		return err, false
	}
	if t.Token == "" {
		return nil, false
	}
	return nil, true
}

func CreateToken(name string, token string) error {
	if err := db.Where("name = ?", name).First(&Node{}).Error; err != nil {
		return err
	}
	t := Token{
		Name:         name,
		Token:        token,
		GenerateTime: time.Now(),
		UpdateTime:   time.Now(),
	}
	if err := db.Create(&t).Error; err != nil {
		return err
	}
	return nil
}

func GetToken(name string) (*Token, error) {
	t := Token{}
	if err := db.Where("name = ?", name).First(&t).Error; err != nil {
		return nil, err
	} else {
		return &t, nil
	}
}

func UpdateToken(name string, token string) error {
	err, ok := CheckToken(name)
	if ok {
		if err := db.Where("name = ?", name).First(&Token{}).Updates(
			map[string]interface{}{
				"Name":       name,
				"Token":      token,
				"UpdateTime": time.Now()}).Error; err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return err
	}
}

func ListToken() ([]*Token, error) {
	var list []*Token
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	} else {
		return list, err
	}
}
