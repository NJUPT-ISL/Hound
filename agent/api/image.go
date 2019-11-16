package api

import (
	"github.com/NJUPT-ISL/Hound/agent/lib"
	"github.com/gin-gonic/gin"
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
	go lib.ImagesPull(c.PostFormArray("imageName"))
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
	lib.ImagesRemove(imageNames)
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
