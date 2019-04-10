package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Nodes struct {
	HostName string `gorm:"unique;not null"`
	Role string `gorm:"size:255"`
	JoinTime time.Time
}

func NodesCreate(db * gorm.DB, hostname string, role string) (string,error){
	node := Nodes{
		HostName:hostname,
		Role:role,
		JoinTime:time.Now(),
	}
	if err := db.Create(&node).Error;err != nil {
		return "0", err
	}
	return node.HostName, nil
}

func CheckNode(db * gorm.DB, hostname string){

}