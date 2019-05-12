package main

import  (
	"./routers"
	"./models"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)


func main() {
	// Init Databases
	models.Setup()
	// Disable Debug mode
	//gin.SetMode(gin.ReleaseMode)
	// Enable Logs
	gin.DisableConsoleColor()
	f, err := os.Create("log/master.log")
	if err != nil{
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := routers.InitRouter()

	_ = r.RunTLS(":8080","pem/server.crt","pem/server.key") // listen and serve on 0.0.0.0:8080
}
