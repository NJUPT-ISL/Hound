package routers

import (
	"../api"
	"github.com/gin-gonic/gin"
)
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	imageGroup := router.Group("/image")
	{
		// Get Action
		imageGroup.GET("/list", api.GetImageList)
		imageGroup.GET("/info", api.GetDockerInfo)
		imageGroup.GET("/prune",api.GetImagePrune)
		// Post Action
		imageGroup.POST("/pull", api.PostImagePull)
		imageGroup.POST("/remove", api.PostImageRemove)
	}
	return router
}