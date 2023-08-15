package main

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

	r.Run("192.168.195.1:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
