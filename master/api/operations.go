package api

import (
	"github.com/NJUPT-ISL/Hound/master/models"
	"github.com/NJUPT-ISL/Hound/master/operations"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
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

	for _, n := range nodes {
		go func() {
			operations.Prune(n)
		}()
	}
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
	var stop <-chan struct{}
	pieces := len(nodes)
	toProcess := make(chan string, pieces)
	for _, n := range nodes {
		toProcess <- n
	}
	close(toProcess)
	if pieces < workers {
		workers = pieces
	}
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for n := range toProcess {
				select {
				case <-stop:
					return
				default:
					operations.Pull(n, c.PostFormArray("imageName"))
				}
			}
		}()
	}
	wg.Wait()
	workers = 10
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
	for _, n := range nodes {
		go func() {
			operations.Remove(n, c.PostFormArray("imageName"))
		}()
	}
	c.JSON(200, gin.H{
		"state": "ok",
	})
}
