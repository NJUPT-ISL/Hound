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
		nodesGroup.POST("join",api.PostNodeJoin)
		nodesGroup.GET("list",api.GetNodeList)
	}
	tokensGroup := router.Group("/tokens")
	{
		tokensGroup.POST("send",api.SendToken)
		tokensGroup.GET("list",api.GetTokenList)
	}
	return router
}
