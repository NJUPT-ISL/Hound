package models

import (
	"github.com/jinzhu/gorm"
)

type Label struct {
	gorm.Model
	Name  string
	Nodes []Node `gorm:"many2many:label_nodes;"`
}

func LabelCreate(name string) error {
	L := Label{
		Name: name,
	}
	if err := db.Create(&L).Error; err != nil {
		return err
	}
	return nil
}

func LabelCheck(name string) (error, bool) {
	if err := db.Where("name = ?", name).First(&Label{}).Error; err != nil {
		return err, false
	}
	return nil, true
}

func LabelQuery(name string) ([]*Label, error) {
	var labels []*Label
	if err := db.Where("name = ?", name).Find(&labels).Error; err != nil {
		return nil, err
	} else {
		return labels, nil
	}
}

func LabelDelete(name string) error {
	var labels []*Label
	err, ok := LabelCheck(name)
	if ok {
		if err := db.Where("name = ?", name).Find(&labels).Delete(&labels).Error; err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return err
	}
}

func LabelListAll() ([]*Label, error) {
	var labelList []*Label
	if err := db.Find(&labelList).Error; err != nil {
		return nil, err
	} else {
		return labelList, nil
	}
}

func LabelOnlyList() ([]*Label, error) {
	var list []*Label
	if err := db.Table("labels").Select("distinct label").Scan(&list).Error; err != nil {
		return nil, err
	} else {
		print(list)
		return list, nil
	}
}
