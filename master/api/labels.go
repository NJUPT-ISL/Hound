package api

import (
	"../models"
	"github.com/gin-gonic/gin"
	"log"
)

func GetLabelList(c *gin.Context){
	if label,err := models.LabelListAll();err !=nil {
		c.JSON(200,"failed")
	}else {
		c.JSON(200,label)
	}
}

func PostAddLabel(c *gin.Context){
	node := c.PostForm("node")
	label := c.PostForm("label")
	if err := models.LabelCreate(node,label);err != nil {
		c.JSON(200,"failed")

	}else {
		c.JSON(200,"ok")
		log.Printf("Label added, node: "+node+" label: "+label)
	}
}