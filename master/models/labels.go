package models

type Labels struct {
	HostName string `gorm:"unique;not null"`
	Label string
}

func LabelCreate(hostname string, labelname string) error{
	label := Labels{
		HostName:hostname,
		Label:labelname,
	}
	if err := db.Create(&label).Error;err != nil {
		return err
	}
	return nil
}

func LabelCheck(labelname string) (error,bool){
	if err := db.Where("label = ?",labelname).First(&Labels{}).Error;err != nil {
		return err,false
	}
	return nil,true
}

func LabelQuery(labelname string)([] *Labels,error){
	var labels [] *Labels
	if err := db.Where("label = ?",labelname).Find(&labels).Error;err != nil{
		return nil,err
	} else {
		return labels, nil
	}
}

