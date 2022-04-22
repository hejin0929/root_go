package main

import (
	"github.com/gin-gonic/gin"
	"modTest/module"
	"modTest/utlis/token"
	"net/http"
	"strings"
)

// UserVerify
// 统一处理用户验证
func UserVerify() gin.HandlerFunc {
	return func(g *gin.Context) {
		var url = g.Request.URL.Path

		// 过滤静态资源
		if strings.Index(url, "oss") != -1 || strings.Index(url, "ws") != -1 || strings.Index(url, "swagger") != -1 || strings.Index(url, "login") != -1 || strings.Index(url, "phone_code") != -1 || strings.Index(url, "upload") != -1 {
			return
		}

		// token验证
		_, err := token.VerifyToken(g)
		if err != nil {
			g.JSON(http.StatusOK, module.Resp{MgsCode: 500, MgsText: "token失败"})
			g.Abort()
		}
	}
}
