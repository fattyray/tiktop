package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktop/entity"
)

func Jwt2r() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		if token == "" {
			token = c.Query("token")
		}
		if token == "" {
			c.JSON(http.StatusBadRequest, entity.Response{
				StatusCode: 1, StatusMsg: "Not login",
			})
			c.Abort()
		}
		_, errToken := Gettoken(token)
		if errToken != nil {
			c.JSON(http.StatusBadRequest, entity.Response{
				StatusCode: 1, StatusMsg: errToken.Error(),
			})
			c.Abort()
		}

		c.Next()
	}
}
