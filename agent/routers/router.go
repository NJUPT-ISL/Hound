package routers

import (
	"github.com/gin-gonic/gin"
	"../api"
)
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	imageGroup := router.Group("/image")
	{
		imageGroup.GET("/list", api.GetImageList)
		imageGroup.POST("/pull", api.PostImagePull)
		imageGroup.POST("/remove", api.PostImageRemove)
	}
	return router
}