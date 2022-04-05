package home

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"modTest/utlis/token"
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
	//authData := g.Request.Header.Get("Authorization") // 获取token

	action, err := token.VerifyToken(g)

	if err != nil {

		g.JSON(200, gin.H{"err": err.Error()})
		return
	}

	fmt.Println("this is a ?? ", action)
}
