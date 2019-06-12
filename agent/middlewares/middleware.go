package middlewares

import (
	"../tokens"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func TokenRequestMiddleware(T *tokens.Token) gin.HandlerFunc {
	return func(c *gin.Context) {
		if key := c.GetHeader("Hound_Key");key != "" {
			if key == os.Getenv("Hound_Key")  {
				log.Printf("Get Hound Key: "+key)
				c.JSON(200,gin.H{
					"token":T.GetToken(),
				})
				c.Abort()
				return
			}else {
				log.Printf("Get Hound Key: "+key)
				c.JSON(401,gin.H{
					"message":"Authentication failed",
				})
				c.Abort()
				return
			}
		}else {
			c.Next()
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
		if action := c.GetHeader("refresh");action != ""{
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
		}else {
			c.Next()
		}
	}
}
