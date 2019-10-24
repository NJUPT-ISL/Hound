package api

import (
	"github.com/NJUPT-ISL/Hound/master/models"
	"github.com/gin-gonic/gin"
	"log"
)

func GetToken(c *gin.Context) {
	Hostname := c.PostForm("Host")
	Token := c.PostForm("Token")
	if _, ok := models.CheckToken(Hostname); !ok {
		if err := models.CreateToken(Hostname, Token); err != nil {
			c.JSON(200, gin.H{
				"state": "failed",
			})
			log.Println(err)
		} else {
			c.JSON(200, gin.H{
				"state": "ok",
			})
			log.Printf("Get Node Token,HostName: " + Hostname + " Token:" + Token)
		}
	} else {

		if err := models.UpdateToken(Hostname, Token); err != nil {
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
	list, err := models.ListToken()
	if err != nil {
		c.JSON(200, gin.H{
			"state": "failed",
		})
	} else {
		c.JSON(200, list)
	}
}
