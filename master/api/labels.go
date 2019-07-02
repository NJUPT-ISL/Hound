package api

import (
	"../models"
	"github.com/gin-gonic/gin"
	"log"
)

func GetLabelListAll(c *gin.Context){
	if label,err := models.LabelListAll();err !=nil {
		c.JSON(200,gin.H{
			"state":"failed",
		})
	}else {
		c.JSON(200,label)
	}
}

func GetLabelOnlyList(c *gin.Context){
	if label,err := models.LabelOnlyList();err !=nil {
		c.JSON(200,gin.H{
			"state":"failed",
		})
	}else {
		c.JSON(200,label)
	}
}

func PostAddLabel(c *gin.Context){
	node := c.PostForm("node")
	label := c.PostForm("label")
	if err := models.LabelCreate(node,label);err != nil {
		c.JSON(200,gin.H{
			"state":"failed",
		})

	}else {
		c.JSON(200,gin.H{
			"state":"ok",
		})
		log.Printf("Label added, node: "+node+" label: "+label)
	}
}

func PostLabelNodelist(c *gin.Context){
	label := c.PostForm("label")
	if nodes,err := models.NodeLabelList(label); err != nil{
		c.JSON(200,gin.H{
			"state":"failed",
		})
	} else {
		c.JSON(200,nodes)
	}
}