package api

import (
	"github.com/NJUPT-ISL/Hound/master/log"
	"github.com/NJUPT-ISL/Hound/master/models"
	"github.com/gin-gonic/gin"
)

func GetLabelListAll(c *gin.Context) {
	if label, err := models.ListAllLabels(); err != nil {
		log.ErrPrint(err)
		c.JSON(200, gin.H{
			"state": "failed",
		})
	} else {
		c.JSON(200, label)
	}
}

func GetLabelOnlyList(c *gin.Context) {
	if label, err := models.OnlyListLabels(); err != nil {
		log.ErrPrint(err)
		c.JSON(200, gin.H{
			"state": "failed",
		})
	} else {
		c.JSON(200, label)
	}
}

func PostAddLabelToNode(c *gin.Context) {
	node := c.PostForm("node")
	label := c.PostForm("label")
	if err := models.CreateLabelWithNode(label, node); err != nil {
		c.JSON(200, gin.H{
			"state": "failed",
		})
		log.ErrPrint(err)
		return
	}
	c.JSON(200, gin.H{
		"state": "ok",
	})
	log.Print("Label added, node: " + node + " label: " + label)
}

func PostAddLabelToNodes(c *gin.Context) {
	nodes := c.PostFormArray("node")
	label := c.PostForm("label")
	for _, node := range nodes {
		if err := models.CreateLabelWithNode(label, node); err != nil {
			c.JSON(200, gin.H{
				"state": "failed",
			})
			log.ErrPrint(err)
		} else {
			c.JSON(200, gin.H{
				"state": "ok",
			})
			log.Print("Label added, node: " + node + " label: " + label)
		}
	}
}

func PostLabelNodelist(c *gin.Context) {
	label := c.PostForm("label")
	if nodes, err := models.GetLabelNodes(label); err != nil {
		log.ErrPrint(err)
		c.JSON(200, gin.H{
			"state": "failed",
		})
	} else {
		c.JSON(200, nodes)
	}
	c.JSON(200, label)
}
