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
// @Param token header  string true "token"
// @Success 200 {object} home.UserMessage true "JSON数据"
// @Failure      400  {object}  module.HttpErrs
// @Failure      404  {object}  module.HttpErrs
// @Failure      500  {object}  module.HttpErrs
// @Router /api/chum/search/{phone} [get]
func SearchUserPaths(g *gin.Context) {
	chum.SearchUser(g)
}
