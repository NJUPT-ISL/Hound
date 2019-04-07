package main
import  (
	"./routers"
)


func main() {
	r := routers.InitRouter()
	_ = r.RunTLS(":8080","pem/server.crt","pem/server.key") // listen and serve on 0.0.0.0:8080
}
