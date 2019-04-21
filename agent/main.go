package main

import (
	"github.com/robfig/cron"
	"./lib"
	"./routers"
	"./tokens"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)


func main() {

	// Set cron
	cronJob := cron.New()
	spec := "*/5 * * * * ?"
	if err := cronJob.AddFunc(spec, func() {
		lib.SendUpdate()
	}) ; err != nil {
		panic(err)
	}

	lib.SendJoin()
	cronJob.Start()
	// Disable Debug mode
	//gin.SetMode(gin.ReleaseMode)
	// Enable Logs
	gin.DisableConsoleColor()
	// GenerateToken
	token := tokens.Token{}
	token.GenerateToken()
	log.Printf("The Hound Agent token is:"+token.GetToken())
	//Send Token
	lib.SendToken(token.GetToken())
	f, err := os.Create("log/agent.log")
	if err != nil{
		panic(err)
	}
	Addr := ":8081"
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.Printf("Hound Service Agent is running at"+Addr)
	r := routers.InitRouter(&token)
	_ = r.RunTLS(Addr,"pem/server.crt","pem/server.key") // listen and serve on 0.0.0.0:8080
}
