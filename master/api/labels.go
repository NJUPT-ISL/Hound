package api

import (
	"github.com/NJUPT-ISL/Hound/master/models"
	"github.com/gin-gonic/gin"
	"log"
)

func GetLabelListAll(c *gin.Context) {
	if label, err := models.ListAllLabels(); err != nil {
		c.JSON(200, gin.H{
			"state": "failed",
		})
	} else {
		c.JSON(200, label)
	}
}

func GetLabelOnlyList(c *gin.Context) {
	if label, err := models.OnlyListLabels(); err != nil {
		c.JSON(200, gin.H{
			"state": "failed",
		})
	} else {
		c.JSON(200, label)
	}
}

func PostAddLabel(c *gin.Context) {
	node := c.PostForm("node")
	label := c.PostForm("label")
	if err := models.CreateLabelWithNode(label, node); err != nil {
		c.JSON(200, gin.H{
			"state": "failed",
		})
	} else {
		c.JSON(200, gin.H{
			"state": "ok",
		})
		log.Printf("Label added, node: " + node + " label: " + label)
	}
	c.JSON(200, node+label)
}

func PostLabelNodelist(c *gin.Context) {
	label := c.PostForm("label")
	if err, nodes := models.GetLabelNodes(label); err != nil {
		c.JSON(200, gin.H{
			"state": "failed",
		})
	} else {
		c.JSON(200, nodes)
	}
	c.JSON(200, label)
}
