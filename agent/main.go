package main
import  (
	"./routers"
)


func main() {
	r := routers.InitRouter()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
