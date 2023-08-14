package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"publish_list": "OK",
	})
}
