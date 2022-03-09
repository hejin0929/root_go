package main

import (
	"github.com/gin-gonic/gin"
	"modTest/utlis/token"
	"strings"
)

// UserVerify
// 统一处理用户验证
func UserVerify() gin.HandlerFunc {
	return func(g *gin.Context) {
		var url = g.Request.URL.Path

		// 过滤静态资源
		if strings.Index(url, "oss") != -1 {
			return
		}

		// 放行swagger文档
		if strings.Index(url, "swagger") != -1 {
			return
		}

		// 放行登陆 注册 验证码接口
		if strings.Index(url, "login") != -1 || strings.Index(url, "phone_code") != -1 {

			return
		}
		// token验证
		token.VerifyToken(g)
		return
	}
}
