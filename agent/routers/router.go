package routers

import (
	"../api"
	"github.com/gin-gonic/gin"
	"../middlewares"
	"../tokens"
)

func InitRouter(token *tokens.Token) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	// Enable Token
	router.Use(middlewares.TokenAuthMiddleware(token))
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
