package loginPaths

import (
	"github.com/gin-gonic/gin"
	"modTest/component/login"
)

// LoginNewsUserPaths
// 注册用户处理逻辑
// @Tags Login
// @Summary 用户注册操作
// @ID LoginNewsUserPaths
// @Param data body	login.UserSignType true "JSON数据"
// @Success 200 {object} login.UserSignReps true "JSON数据"
// @Router /api/login/user/sign [post]
func LoginNewsUserPaths(g *gin.Context) {
	//user := new(login2.UserName)
	//if g.Bind(user) != nil {
	//	g.JSON(http.StatusUnprocessableEntity, "参数错误")
	//}
	login.SignUser(g)

}

// LoginInPasswordPaths
// 用户账号密码登录
// @Tags Login
// @Summary 用户登录操作
// @ID LoginInPasswordPaths
// @Param data body	login.UserName true "JSON数据"
// @Success 200 {array} login.UserBody true "JSON数据"
// @Router /api/login/user/login [post]
func LoginInPasswordPaths(g *gin.Context) {
	//params := new(login2.UserName)
	//
	//if g.Bind(params) != nil {
	//	g.JSON(http.StatusUnprocessableEntity, "参数错误")
	//	return
	//}

	login.LoginsUserPassword(g)
}

// LoginInCodePaths
// 用户手机验证码登录
// @Tags Login
// @Summary 用户登录操作
// @ID LoginsUserCode
// @Param data body	login.UserCode true "JSON数据"
// @Success 200 {array} login.UserBody true "JSON数据"
// @Router /api/login/user/login_code [post]
func LoginInCodePaths(g *gin.Context) {
	//params := new(login2.UserCode)
	//if g.Bind(params) != nil {
	//	g.JSON(http.StatusUnprocessableEntity, "参数错误")
	//	return
	//}
	login.LoginsUserCode(g)
}
