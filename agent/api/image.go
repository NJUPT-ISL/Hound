package api

import (
	"fmt"
	"github.com/NJUPT-ISL/Hound/agent/lib"
	"github.com/gin-gonic/gin"
	"log"
)

// Get Action
func GetImageList(c *gin.Context) {
	images, err := lib.ListAllImages()
	if err != nil {
		panic(err)
	}
	c.JSON(200, images)
}

func GetImagePrune(c *gin.Context) {
	report, err := lib.ImagesPrune()
	if err != nil {
		panic(err)
	}
	c.JSON(200, report)

}

func GetDockerInfo(c *gin.Context) {
	info, err := lib.DockerInfo()
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"Images":            info.Images,
		"SystemTime":        info.SystemTime,
		"KernelVersion":     info.KernelVersion,
		"OperatingSystem":   info.OperatingSystem,
		"NCPU":              info.NCPU,
		"DockerVersion":     info.ServerVersion,
		"ContainersRunning": info.ContainersRunning,
		"ContainersPaused":  info.ContainersPaused,
		"ContainersStopped": info.ContainersStopped,
	})
}

// PostAction
func PostImagePull(c *gin.Context) {
	imageNames := c.PostFormArray("imageName")
	go func() {
		for _, imageName := range imageNames {

			go func() {
				_, err := lib.ImagePull(imageName)
				if err != nil {
					log.Println(err)
				}
			}()

		}
	}()
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func PostImageRemove(c *gin.Context) {

	imageNames := c.PostFormArray("imageName")
	force := false
	if c.PostForm("Force") == "true" {
		force = true
	}
	go func() {
		for _, imageName := range imageNames {
			go func() {
				_, err := lib.ImageRemove(imageName, force)
				if err != nil {
					fmt.Println(err)
				}
			}()
		}
	}()
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
