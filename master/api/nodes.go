package api

import (
	"../models"
	"github.com/gin-gonic/gin"
	"log"
)


func PostJoin(c *gin.Context){
	Hostname := c.PostForm("Host")
	Role := c.PostForm("Role")
	if err := models.NodesCreate(Hostname,Role ); err != nil {
		c.JSON(200,gin.H{
			"state": "failed",
		})
	} else {
		c.JSON(200,gin.H{
			"state": "ok",
		})
		log.Printf(Hostname+Role)
	}
}

func GetList(c *gin.Context){
	list,err := models.NodeList()
	if err != nil{
		c.JSON(200,gin.H{
			"message":"failed",
		})
	} else {
		c.JSON(200, list)
	}
}