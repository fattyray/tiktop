package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Jwt2r() gin.HandlerFunc {
	return func(x *gin.Context) {
		var code int
		code = 200
		token := x.Query("token")

		if token == "" {
			code = 401
		} else {
			claims, err := Gettoken(token)
			if err != nil {
				code = 402
			} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
				code = 403
			}
		}
		if code != 200 {
			x.JSON(http.StatusUnauthorized, http.Response{StatusCode: code})

			return
		}
		x.Next()
	}
}
