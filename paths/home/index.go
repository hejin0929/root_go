package home

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"modTest/component/home"
	"modTest/module"
	"modTest/utlis/key"
	"net/http"
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

// AppHomeKeyPaths 登陆首页
// @Tags Home
// @Summary 获取数据加密
// @Description  get string by ID
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Param token header  string true "token"
// @Success      200  {object}  home.KeysRes
// @Failure      400  {object}  module.HttpErrs
// @Failure      404  {object}  module.HttpErrs
// @Failure      500  {object}  module.HttpErrs
// @Router       /api/home/key [get]
func AppHomeKeyPaths(g *gin.Context) {
	files := key.GenerateRSAKey(2048)

	private, _ := ioutil.ReadFile(files[0])

	public, _ := ioutil.ReadFile(files[1])

	res := KeysRes{}

	res.Body.Private = string(private)

	res.Body.Public = string(public)

	g.JSON(http.StatusOK, module.ResponseSuccess(res.Body))

}
