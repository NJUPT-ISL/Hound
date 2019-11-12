package routers

import (
	"fmt"
	"github.com/NJUPT-ISL/Hound/master/api"
	dash "github.com/NJUPT-ISL/Hound/master/api/dashboard"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	//router.Use(gin.Logger())

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("[Hound Master]%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
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
	nodesGroup := router.Group("/nodes")
	{
		nodesGroup.POST("join", api.PostNodeJoin)
		nodesGroup.GET("list", api.GetNodeList)
		nodesGroup.POST("update", api.PostNodeUpdate)
	}
	tokensGroup := router.Group("/tokens")
	{
		tokensGroup.POST("send", api.GetToken)
		tokensGroup.GET("list", api.GetTokenList)
	}
	operationsGroup := router.Group("/operations")
	{
		operationsGroup.POST("prune", api.PostNodePrune)
		operationsGroup.POST("pull", api.PostNodePullImage)
		operationsGroup.POST("remove", api.PostNodeRemoveImage)
	}
	labelsGroup := router.Group("/labels")
	{
		labelsGroup.GET("list", api.GetLabelListAll)
		labelsGroup.GET("only", api.GetLabelOnlyList)
		labelsGroup.POST("add", api.PostAddLabelToNode)
		labelsGroup.POST("adds", api.PostAddLabelToNodes)
		labelsGroup.POST("node", api.PostLabelNodelist)
	}
	labelOpsGroup := router.Group("/labelops")
	{
		labelOpsGroup.POST("prune", api.PostLabelPrune)
		labelOpsGroup.POST("pull", api.PostLabelPull)
		labelOpsGroup.POST("remove", api.PostLabelRemove)
	}
	return router
}
