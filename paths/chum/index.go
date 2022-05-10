package chum

import (
	"github.com/gin-gonic/gin"
	"modTest/component/chum"
	"modTest/module"
	chum2 "modTest/types/chum"
	"net/http"
)

// SearchUserPaths
// 获取验证码
// @Tags Chum
// @Summary 索搜好友信息
// @ID SearchUserPaths
// @Produce  json
// @Accept  json
// @Param phone path string true "朋友手机号"
// @Param token header  string true "token"
<<<<<<< HEAD
// @Success 200 {object} chum2.ResChum true "JSON数据"
=======
// @Success 200 {object} chum2.SearchUserRes true "JSON数据"
>>>>>>> 7c5cfbafc3b65ab7e1446ab7bf96cee5a60d0051
// @Failure      400  {object}  module.HttpErrs
// @Failure      404  {object}  module.HttpErrs
// @Failure      500  {object}  module.HttpErrs
// @Router /api/chum/search/{phone} [get]
func SearchUserPaths(g *gin.Context) {

	mobile := g.Query("phone")

	if mobile == "" {
		g.JSON(http.StatusOK, module.ResponseErrorParams("参数不全"))
		return
	}

	user := chum.SearchUser(mobile)

	g.JSON(http.StatusOK, module.ResponseSuccess(user))

	return
}

// AddUserFriendPaths
// 获取验证码
// @Tags Chum
// @Summary 用户验证码登录
// @ID AddUserFriendPaths
// @Produce  json
// @Accept  json
// @Param token header  string true "token"
// @Param data  body  AddReq true "发送数据"
// @Success 200 {object} chum2.AddRes true "JSON数据"
// @Failure      400  {object}  module.HttpErrs
// @Failure      404  {object}  module.HttpErrs
// @Failure      500  {object}  module.HttpErrs
// @Router /api/chum/add [post]
func AddUserFriendPaths(g *gin.Context) {
	reqChum := new(chum2.AddReq)

	if g.Bind(reqChum) != nil {
		g.JSON(http.StatusOK, module.ResponseErrorParams("参数不全"))
		return
	}

	err := chum.AddChumUser(reqChum)

	if err != nil {
		g.JSON(http.StatusOK, module.ResponseServerError("添加失败"))
		return
	}

	g.JSON(http.StatusOK, module.ResponseSuccess("添加成功"))
	return
}
