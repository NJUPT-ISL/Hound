package routers

import (
	"../api"
	"../middlewares"
	"../tokens"
	"github.com/gin-gonic/gin"
)

func InitRouter(token *tokens.Token) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	// Enable Token
	router.Use(middlewares.TokenRequestMiddleware(token),middlewares.TokenAuthMiddleware(token),middlewares.TokenRefreshMiddleware(token))
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
	//tokenGroup := router.Group("/token")
	//{
	//	tokenGroup.GET("/refresh", api.GetRefresh)
	//}
	return router
}
