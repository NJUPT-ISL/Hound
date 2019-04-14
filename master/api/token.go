package api

import (
	"../models"
	"github.com/gin-gonic/gin"
	"log"
)

func SendToken(c *gin.Context){
	Hostname := c.PostForm("Host")
	Token := c.PostForm("Token")
	if err := models.TokenCreate(Hostname,Token); err != nil {
		c.JSON(200,gin.H{
			"state": "failed",
		})
	} else {
		c.JSON(200,gin.H{
			"state": "ok",
		})
		log.Printf("Get Node Token,HostName: "+Hostname+" Token:"+Token)
	}
}

func GetTokenList(c *gin.Context){
	list,err := models.TokenList()
	if err != nil{
		c.JSON(200,gin.H{
			"message":"failed",
		})
	} else {
		c.JSON(200, list)
	}
}