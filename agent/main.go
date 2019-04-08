package main
import  (
	"./routers"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"log"
	"./tokens"
)


func main() {
	// Disable Debug mode
	gin.SetMode(gin.ReleaseMode)
	// Enable Logs
	gin.DisableConsoleColor()
	// GenerateToken
	token := tokens.Token{}
	token.GenerateToken()
	log.Printf("The Hound Agent token is:"+token.GetToken())

	f, err := os.Create("log/agent.log")
	if err != nil{
		panic(err)
	}
	Addr := ":8080"
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.Printf("Hound Service Agent is running at"+Addr)
	r := routers.InitRouter(&token)
	_ = r.RunTLS(Addr,"pem/server.crt","pem/server.key") // listen and serve on 0.0.0.0:8080
}
