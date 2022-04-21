package chum

import (
	"github.com/gin-gonic/gin"
	"modTest/component/chum"
)

// SearchUserPaths
// 获取验证码
// @Tags Chum
// @Summary 用户验证码登录
// @ID SearchUserPaths
// @Produce  json
// @Accept  json
// @Param phone path string true "朋友手机号"
// @Success 200 {object} ResChum true "JSON数据"
// @Router /api/chum/search/{phone} [get]
func SearchUserPaths(g *gin.Context) {
	chum.SearchUser(g)
}
