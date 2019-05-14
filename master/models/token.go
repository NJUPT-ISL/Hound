package models

import "time"

type Tokens struct {
	HostName string `gorm:"unique;not null"`
	Token string
	GenerateTime time.Time
	UpdateTime time.Time
}

func TokenCreate(hostname string, token string) error {
	t := Tokens{
		HostName:hostname,
		Token:token,
		GenerateTime:time.Now(),
		UpdateTime:time.Now(),
	}
	if err := db.Create(&t).Error;err != nil {
		return err
	}
	return nil
}

func TokenCheck(hostname string) (error,bool){
	if err := db.Where("host_name = ?",hostname).First(&Tokens{}).Error;err != nil {
		return err,false
	}
	return nil,true
}

func TokenQuery(hostname string)(*Tokens,error){
	t := Tokens{}
	if err := db.Where("host_name = ?",hostname).First(&t).Error;err != nil{
		return nil,err
	} else {
		return &t, nil
	}
}

func TokenUpdate (hostname string,token string) error {
	err,ok := TokenCheck(hostname)
	if ok {
		if err := db.Where("host_name = ?",hostname).First(&Tokens{}).Updates(
			map[string]interface{}{
				"HostName": hostname,
				"Token": token,
				"UpdateTime": time.Now()}).Error; err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return err
	}
}

func TokenDelete (hostname string) error {
	err,ok := NodeCheck(hostname)
	if ok {
		if err := db.Where("host_name = ?",hostname).First(&Tokens{}).Delete(&Nodes{}).Error; err != nil{
			return err
		} else {
			return nil
		}
	} else {
		return err
	}
}

func TokenList() ([]*Tokens,error) {
	var list [] *Tokens
	if err := db.Find(&list).Error ; err !=nil{
		return nil,err
	} else {
		return list,err
	}
}