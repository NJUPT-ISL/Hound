package api

import (
	"../models"
	"github.com/gin-gonic/gin"
	"log"
)

func SendToken(c *gin.Context) {
	Hostname := c.PostForm("Host")
	Token := c.PostForm("Token")
	if _, ok := models.TokenCheck(Hostname); !ok {
		if err := models.TokenCreate(Hostname, Token); err != nil {
			c.JSON(200, gin.H{
				"state": "failed",
			})
		} else {
			c.JSON(200, gin.H{
				"state": "ok",
			})
			log.Printf("Get Node Token,HostName: " + Hostname + " Token:" + Token)
		}
	} else {

		if err := models.TokenUpdate(Hostname, Token); err != nil {
			c.JSON(200, gin.H{
				"state": "failed",
			})
			panic(err)
		} else {
			c.JSON(200, gin.H{
				"state": "ok",
			})
			log.Printf("Node: " + Hostname + " Token Updated: " + Token)
		}
	}

}

func GetTokenList(c *gin.Context) {
	list, err := models.TokenList()
	if err != nil {
		c.JSON(200, gin.H{
			"state": "failed",
		})
	} else {
		c.JSON(200, list)
	}
}
