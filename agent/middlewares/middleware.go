package middlewares

import (
	"github.com/gin-gonic/gin"
	"../tokens"
	"log"
)

func TokenAuthMiddleware(T *tokens.Token) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		log.Printf("The request token is "+token)
		if token == "" {
			c.JSON(401,gin.H{
				"message":"API token required",
			})
			c.Abort()
			return
		}else if !T.VerifyToken(token){
			c.JSON(403,gin.H{
				"message":"Invalid API token",
			})
			c.Abort()
			return
		}else {
			c.Next()
		}
	}
}