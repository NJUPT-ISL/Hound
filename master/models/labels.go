package models

type Labels struct {
	Node string `gorm:"not null;unique"`
	Label string
}

func LabelCreate(node string, labelName string) error{
	L := Labels{
		Node:node,
		Label:labelName,
	}
	if err := db.Create(&L).Error;err != nil {
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
func LabelDelete (labelname string) error {
	var labels [] *Labels
	err,ok := LabelCheck(labelname)
	if ok {
		if err := db.Where("label = ?",labelname).Find(&labels).Delete(&labels).Error; err != nil{
			return err
		} else {
			return nil
		}
	} else {
		return err
	}
}
func NodeLabelList (labelname string) ([] *string,error){
	var nodeList [] *string
	err,ok := LabelCheck(labelname)
	if ok {
		if err := db.Select("node").Where("label = ?",labelname).Find(&nodeList).Error;err != nil{
			return nil,err
		}else {
			return nodeList,nil
		}
	}else {
		return nil,err
	}
}

func LabelListAll() ([] *Labels,error){
	var labelList [] *Labels
	if err := db.Find(&labelList).Error;err != nil{
		return nil,err
	}else {
		return labelList,nil
	}
}