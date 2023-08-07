package midware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 目的是写一个gin的中间件，然后每个需要token进行用户鉴权的就需要先经过这个中间件来判断
// 不能通过的情况有如下的几种
// 1.没有token
// 2.token出错
// 3.token的session超时
// 4.运行鉴权的函数的时候出现了错误
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
