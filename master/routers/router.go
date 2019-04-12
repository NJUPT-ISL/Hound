package routers

import (
	"github.com/gin-gonic/gin"
	"../api"
	)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	nodesGroup := router.Group("/nodes")
	{
		nodesGroup.POST("join",api.PostJoin)
		nodesGroup.GET("list",api.GetList)
	}
	return router
}
