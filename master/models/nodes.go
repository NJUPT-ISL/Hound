package models

import (
	"time"
)

type Node struct {
	Name            string `gorm:"PRIMARY_KEY;unique;not null"`
	Role            string `gorm:"size:255"`
	KernelVersion   string
	OperatingSystem string
	DockerVersion   string
	JoinTime        time.Time
	UpdateTime      time.Time
}

func CreateNode(name string, role string, kv string, os string, dv string) error {
	node := Node{
		Name:            name,
		Role:            role,
		KernelVersion:   kv,
		OperatingSystem: os,
		DockerVersion:   dv,
		JoinTime:        time.Now(),
		UpdateTime:      time.Now(),
	}
	if err := db.Create(&node).Error; err != nil {
		return err
	}
	return nil
}

func GetNode(name string) (err error, node Node) {
	n := Node{}
	if err := db.Where("name = ?", name).First(&n).Error; err != nil {
		return err, n
	}
	return nil, n
}

func CheckNode(name string) (error, bool) {
	if err := db.Where("name = ?", name).First(&Node{}).Error; err != nil {
		return err, false
	}
	return nil, true
}

func UpdateNode(name string, role string, kv string, os string, dv string) error {
	var node = Node{}
	err, ok := CheckNode(name)
	if ok {
		if err = db.Model(&node).Where("name = ?", name).Updates(
			map[string]interface{}{
				"Role":            role,
				"KernelVersion":   kv,
				"OperatingSystem": os,
				"DockerVersion":   dv,
				"UpdateTime":      time.Now()}).Error; err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return err
	}
}

func ListNodes() ([]*Node, error) {
	var list []*Node
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	} else {
		return list, err
	}
}
