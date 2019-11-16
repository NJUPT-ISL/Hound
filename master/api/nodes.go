package api

import (
	"github.com/NJUPT-ISL/Hound/master/log"
	"github.com/NJUPT-ISL/Hound/master/models"
	"github.com/gin-gonic/gin"
)

func PostNodeJoin(c *gin.Context) {

	if _, ok := models.CheckNode(c.PostForm("Host")); ok != true {
		if err := models.CreateNode(c.PostForm("Host"),
			c.PostForm("Role"),
			c.PostForm("kv"),
			c.PostForm("os"),
			c.PostForm("dv")); err != nil {
			log.Print("Create Node Error: " + c.PostForm("Host"))
			log.ErrPrint(err)
			c.JSON(200, gin.H{
				"state": "failed",
			})
		} else {
			c.JSON(200, gin.H{
				"state": "ok",
			})
			log.Print("New Node joined,HostName: " + c.PostForm("Host") + " Role: " + c.PostForm("Role") + ".")
		}
	} else {
		c.JSON(200, gin.H{
			"state": "ok",
		})
		log.Print("Node " + c.PostForm("Host") + " already exists,and does not need to send a join message.")
	}
}

func PostNodeUpdate(c *gin.Context) {
	if err := models.UpdateNode(
		c.PostForm("Host"),
		c.PostForm("Role"),
		c.PostForm("kv"),
		c.PostForm("os"),
		c.PostForm("dv")); err != nil {
		log.Print("Create Node Error: " + c.PostForm("Host"))
		log.ErrPrint(err)
		c.JSON(200, gin.H{
			"state": "failed",
		})
	} else {
		c.JSON(200, gin.H{
			"state": "ok",
		})
		log.Print("Node: " + c.PostForm("Host") + " Updated, Role:" + c.PostForm("Role"))
	}
}

func GetNodeList(c *gin.Context) {
	list, err := models.ListNodes()
	if err != nil {
		log.ErrPrint(err)
		c.JSON(200, gin.H{
			"state": "failed",
		})
	} else {
		c.JSON(200, list)
	}
}
