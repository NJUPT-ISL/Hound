package main

import (
	"github.com/NJUPT-ISL/Hound/agent/cronjob"
	"github.com/NJUPT-ISL/Hound/agent/lib"
	"github.com/NJUPT-ISL/Hound/agent/routers"
	"github.com/NJUPT-ISL/Hound/agent/tokens"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"io"
	"log"
	"os"
)

func main() {
	lib.SendJoin()
	// Disable Debug mode
	//gin.SetMode(gin.ReleaseMode)
	// Enable Logs
	gin.DisableConsoleColor()
	// GenerateToken
	token := tokens.Token{}
	token.GenerateToken()
	log.Printf("The Hound Agent token is:" + token.GetToken())
	//Send Token
	lib.SendToken(token.GetToken())
	f, err := os.Create("log/agent.log")
	if err != nil {
		panic(err)
	}
	Addr := ":8081"
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.Printf("Hound Service Agent is running at" + Addr)
	//cronjob.SendUpdateJob()
	c := cron.New()
	cronjob.AddDockerPruneJob(c)
	cronjob.AddSendUpdateJob(c)
	c.Start()
	r := routers.InitRouter(&token)
	_ = r.RunTLS(Addr, "pem/server.crt", "pem/server.key") // listen and serve on 0.0.0.0:8081
}
