package main

//parparing to merge and upload to code1024
import (
	"github.com/gin-gonic/gin"
	"os"
	"tiktop/initialize"
)

func main() {
	err := initialize.InitDB()
	if err != nil {
		os.Exit(-1)
	}
	r := gin.Default()

	initRouter(r)

	//r.Run()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
