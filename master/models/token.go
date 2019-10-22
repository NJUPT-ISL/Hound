package models

import "time"

type Token struct {
	Name         string `gorm:"unique;not null"`
	Token        string
	GenerateTime time.Time
	UpdateTime   time.Time
}

func TokenCreate(name string, token string) error {
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

func TokenCheck(name string) (error, bool) {
	if err := db.Where("name = ?", name).First(&Token{}).Error; err != nil {
		return err, false
	}
	return nil, true
}

func TokenQuery(name string) (*Token, error) {
	t := Token{}
	if err := db.Where("name = ?", name).First(&t).Error; err != nil {
		return nil, err
	} else {
		return &t, nil
	}
}

func TokenUpdate(name string, token string) error {
	err, ok := TokenCheck(name)
	if ok {
		if err := db.Where("name = ?", name).First(&Token{}).Updates(
			map[string]interface{}{
				"HostName":   name,
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

func TokenList() ([]*Token, error) {
	var list []*Token
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	} else {
		return list, err
	}
}
