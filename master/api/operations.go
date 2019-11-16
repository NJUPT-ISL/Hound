package api

import (
	"github.com/NJUPT-ISL/Hound/master/models"
	"github.com/NJUPT-ISL/Hound/master/operations"
	"github.com/gin-gonic/gin"
	"log"
)

var workers = 10

func PostNodePrune(c *gin.Context) {
	operations.Prune(c.PostForm("NodeName"))
	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func PostNodePullImage(c *gin.Context) {
	operations.Pull(c.PostForm("NodeName"), c.PostFormArray("imageName"))
	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func PostNodeRemoveImage(c *gin.Context) {
	operations.Remove(c.PostForm("NodeName"), c.PostFormArray("imageName"))
	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func PostLabelPrune(c *gin.Context) {
	nodes, err := models.GetLabelNodes(c.PostForm("Label"))
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"state": "failed",
		})
		return
	}

	operations.ParallelizeWithString(workers, nodes, operations.Prune)
	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func PostLabelPull(c *gin.Context) {
	nodes, err := models.GetLabelNodes(c.PostForm("Label"))
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"state": "failed",
		})
		return
	}
	operations.Parallelize(workers, nodes, c.PostFormArray("imageName"), operations.Pull)
	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func PostLabelRemove(c *gin.Context) {
	nodes, err := models.GetLabelNodes(c.PostForm("Label"))
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"state": "failed",
		})
		return
	}
	operations.Parallelize(workers, nodes, c.PostFormArray("imageName"), operations.Remove)
	c.JSON(200, gin.H{
		"state": "ok",
	})
}
