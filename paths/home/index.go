package home

import (
	"github.com/gin-gonic/gin"
	"modTest/component/home"
)

// AppHomePaths 登陆首页
// @Tags Home
// @Summary 首页接口
// @Description  get string by ID
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Param Authorization header  string true "token"
// @Success      200  {object}  string
// @Failure      400  {object}  module.HttpErrs
// @Failure      404  {object}  module.HttpErrs
// @Failure      500  {object}  module.HttpErrs
// @Router       /api/home/message [get]
func AppHomePaths(g *gin.Context) {
	home.AppHone(g)
}
