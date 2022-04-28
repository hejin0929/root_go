package home

import (
	"github.com/gin-gonic/gin"
	"modTest/component/home"
)

// GetUserMessagePaths 登陆首页
// @Tags User
// @Summary 获取个人信息
// @Description  get string by ID
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Account ID"
// @Param Authorization header  string true "token"
// @Success      200  {object}  home.MessageRes
// @Failure      400  {object}  module.HttpErrs
// @Failure      404  {object}  module.HttpErrs
// @Failure      500  {object}  module.HttpErrs
// @Router       /api/user/user_message/{id} [get]
func GetUserMessagePaths(g *gin.Context) {
	home.GetUserMessage(g)
}

// UpdateUserMessagePaths 修改个人信息
// @Tags User
// @Summary 修改个人信息
// @Description  get string by ID
// @Accept       json
// @Produce      json
// @Param        data   body      home.UserMessage  true  "修改的数据"
// @Param token header  string true "token"
// @Success      200  {object}  home.MessageUpdateRes
// @Failure      400  {object}  module.HttpErrs
// @Failure      404  {object}  module.HttpErrs
// @Failure      500  {object}  module.HttpErrs
// @Router       /api/user/user_message/update [post]
func UpdateUserMessagePaths(g *gin.Context) {
	home.UpdateUserMessage(g)
}
