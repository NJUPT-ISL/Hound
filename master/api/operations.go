package api

import (
	"../operations"
	"github.com/gin-gonic/gin"
)

func PostNodePrune(c *gin.Context){
	operations.Prune(c.PostForm("NodeName"))
	c.JSON(200,gin.H{
		"state": "ok",
	})
}

func PostNodePullImage(c *gin.Context){
	operations.Pull(c.PostForm("NodeName"), c.PostFormArray("imageName"))
	c.JSON(200,gin.H{
		"state": "ok",
	})
}

func PostNodeRemoveImage(c *gin.Context){
	operations.Remove(c.PostForm("NodeName"), c.PostFormArray("imageName"))
	c.JSON(200,gin.H{
		"state": "ok",
	})
}
