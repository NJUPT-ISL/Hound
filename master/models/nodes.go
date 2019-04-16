package models

import (
	"time"
)

type Nodes struct {
	HostName string `gorm:"PRIMARY_KEY;unique;not null"`
	Role string `gorm:"size:255"`
	KernelVersion string
	OperatingSystem string
	DockerVersion string
	JoinTime time.Time
}

func NodesCreate(hostname string, role string, kv string, os string,dv string) error{
	node := Nodes{
		HostName:hostname,
		Role:role,
		KernelVersion: kv,
		OperatingSystem: os,
		DockerVersion: dv,
		JoinTime:time.Now(),
	}
	if err := db.Create(&node).Error;err != nil {
		return err
	}
	return nil
}

func NodeCheck(hostname string) (error,bool){
	if err := db.Where("hostname = ?",hostname).First(&Nodes{}).Error;err != nil {
		return err,true
	}
	return nil,false
}

func NodeQuery(hostname string)(*Nodes,error){
	node := Nodes{}
	if err := db.Where("hostname = ?",hostname).First(&node).Error;err != nil{
		return nil,err
	} else {
		return &node, nil
	}
}

func NodesUpdate (hostname string,role string, kv string, os string,dv string) error {
	err,ok := NodeCheck(hostname)
	if ok {
		if err := db.Where("hostname = ?",hostname).First(&Nodes{}).Update(hostname,role,kv,os,dv,time.Now()).Error; err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return err
	}
}

func NodeDelete (hostname string) error {
	err,ok := NodeCheck(hostname)
	if ok {
		if err := db.Where("hostname = ?",hostname).First(&Nodes{}).Delete(&Nodes{}).Error; err != nil{
			return err
		} else {
			return nil
		}
	} else {
		return err
	}
}

func NodeList () ([]*Nodes,error) {
	var list [] *Nodes
	if err := db.Find(&list).Error ; err !=nil{
		return nil,err
	} else {
		return list,err
	}
}