package home

import (
	"github.com/gin-gonic/gin"
	"modTest/component/home"
)

// GetUserMessagePaths 登陆首页
// @Tags User
// @Summary 首页接口
// @Description  get string by ID
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Account ID"
// @Param Authorization header  string true "token"
// @Success      200  {object}  home.UserMessage
// @Failure      400  {object}  module.HttpErrs
// @Failure      404  {object}  module.HttpErrs
// @Failure      500  {object}  module.HttpErrs
// @Router       /api/user/user_message/{id} [get]
func GetUserMessagePaths(g *gin.Context) {
	home.GetUserMessage(g)
}
