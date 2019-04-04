package api

import (
	"../lib"
	"github.com/gin-gonic/gin"
	"log"
)


func GetImageList(c *gin.Context){
	images, err := lib.ListAllImages()
	if err != nil {
		panic(err)
	}
	c.JSON(200, images)
}

func PostImagePull(c *gin.Context){
	imageName := c.PostForm("imageName")
	go func (){
		_, err := lib.ImagePull(imageName)
		if err != nil {
		panic(err)
		}
	}()
	c.JSON(200,gin.H{
		"message":"OK",
	})
}

func PostImageRemove(c *gin.Context){

	imageName := c.PostForm("imageName")
	force := false
	if c.PostForm("Force") == "true" {
		force = true
	}
	go func (){
		_, err := lib.ImageRemove(imageName, force)
		if err != nil {
			log.Panic(err)
		}
	}()
	c.JSON(200,gin.H{
		"message":"OK",
	})
}
