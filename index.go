package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// UserVerify
// 统一处理用户验证
func UserVerify() gin.HandlerFunc {
	return func(g *gin.Context) {
		var url = g.Request.URL.Path
		//fmt.Println("this is ？？", strings.Index(url, "login"))
		if strings.Index(url, "login") == -1 {
			fmt.Println("待写验证逻辑")
			return
		}

	}
}
