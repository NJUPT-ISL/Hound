package middlewares

import (
	"github.com/gin-gonic/gin"
	"../tokens"
	"log"
	"os"
)

func TokenRequestMiddleware(T *tokens.Token) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("Hound_Key")
		log.Printf(key)
		if key == os.Getenv("Hound_Key")  {
			c.JSON(200,gin.H{
				"token":T.GetToken(),
			})
			c.Abort()
			return
		}else {
			c.JSON(401,gin.H{
				"message":"Authentication failed",
			})
			c.Abort()
			return
		}
	}
}

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

func TokenRefreshMiddleware(T *tokens.Token) gin.HandlerFunc {
	return func(c *gin.Context) {
		action := c.GetHeader("refresh")
		if action == "true" {
			T.RefreshToken()
			c.JSON(200,gin.H{
				"token": T.GetToken(),
			})
			log.Printf("The new token is: "+T.GetToken())
			c.Abort()
			return
		}else {
			c.Next()
		}
	}
}
