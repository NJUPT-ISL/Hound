package routers

import (
	"../api"
	"../middlewares"
	"../tokens"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitRouter(token *tokens.Token) *gin.Engine {
	router := gin.New()
	//router.Use(gin.Logger())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("[Hound Agent]%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	// Enable Token
	router.Use(
		middlewares.TokenRequestMiddleware(token),
		middlewares.TokenAuthMiddleware(token),
		middlewares.TokenRefreshMiddleware(token))
	imageGroup := router.Group("/image")
	{
		// Get Action
		imageGroup.GET("list", api.GetImageList)
		imageGroup.GET("info", api.GetDockerInfo)
		imageGroup.GET("prune",api.GetImagePrune)
		// Post Action
		imageGroup.POST("pull", api.PostImagePull)
		imageGroup.POST("remove", api.PostImageRemove)
	}
	//tokenGroup := router.Group("/token")
	//{
	//	tokenGroup.GET("/refresh", api.GetRefresh)
	//}
	return router
}
