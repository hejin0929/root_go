package loginPaths

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginPaths
// 注册用户处理逻辑
// @Tags Login
// @Summary 用户登录操作
// @ID SignUser
// @Param data body	UserSignType true "JSON数据"
// @Success 200 {object} UserSignReps true "JSON数据"
// @Router /login/user/sign [post]
func LoginPaths(g *gin.Context) {
	user := new(UserName)
	if g.Bind(user) != nil {
		g.JSON(http.StatusUnprocessableEntity, "参数错误")
	}

}

type UserSignType struct {
}

type UserSignReps struct {
}
