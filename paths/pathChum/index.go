package pathChum

import (
	"github.com/gin-gonic/gin"
	"modTest/component/chum"
	"modTest/component/home"
	"modTest/module"
	"modTest/types/typeChum"
	token2 "modTest/utlis/token"
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
// @Success 200 {object} typeChum.ResChum true "JSON数据"
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

	g.JSON(http.StatusOK, module.ResponseSuccess(user.Body))

	return
}

// AddUserFriendPaths
// 添加好友接口
// @Tags Chum
// @Summary 添加好友接口
// @ID AddUserFriendPaths
// @Produce  json
// @Accept  json
// @Param token header  string true "token"
// @Param data  body  typeChum.AddReq true "发送数据"
// @Success 200 {object} typeChum.AddRes true "JSON数据"
// @Failure      400  {object}  module.HttpErrs
// @Failure      404  {object}  module.HttpErrs
// @Failure      500  {object}  module.HttpErrs
// @Router /api/chum/add [post]
func AddUserFriendPaths(g *gin.Context) {

	reqChum := new(typeChum.AddReq)

	if g.Bind(&reqChum) != nil {
		g.JSON(http.StatusOK, module.ResponseErrorParams("参数不全"))
		return
	}

	err := chum.AddChumUser(reqChum)

	if err != nil {
		if err.Error() == "" {
			g.JSON(http.StatusOK, module.ResponseServerError("添加失败"))
			return
		}

		if err.Error() == "添加次数已限制" {
			g.JSON(http.StatusOK, module.ResponseSuccess(err.Error()))
			return
		}

		g.JSON(http.StatusOK, module.ResponseServerError(err.Error()))
		return
	}

	g.JSON(http.StatusOK, module.ResponseSuccess("添加成功"))
	return
}

// GetChumApply
// 获取验证码
// @Tags Chum
// @Summary 获取好友申请
// @ID getChumApply
// @Produce  json
// @Accept  json
// @Param token header  string true "token"
// @Param id  path  string true "发送数据"
// @Success 200 {object} typeChum.ApplyUserRes true "JSON数据"
// @Failure      400  {object}  module.HttpErrs
// @Failure      404  {object}  module.HttpErrs
// @Failure      500  {object}  module.HttpErrs
// @Router /api/chum/apply [get]
func GetChumApply(g *gin.Context) {
	id := g.Query("id")

	data := chum.GetApplyUser(id)

	if len(data) == 0 {
		g.JSON(http.StatusOK, module.ResponseSuccess(make([]int, 0)))
		return
	}

	g.JSON(http.StatusOK, module.ResponseSuccess(&data))

}

// ApplyUpdate 更新好友请求 ApplyUser
// @Tags Chum
// @Summary 同意好友申请
// @ID ApplyUpdate
// @Produce  json
// @Accept  json
// @Param token header  string true "token"
// @Param data  body  typeChum.ApplyUpdateParams true "发送数据"
// @Success 200 {object} module.ResponseBodyInString true "JSON数据"
// @Failure      400  {object}  module.HttpErrs
// @Failure      404  {object}  module.HttpErrs
// @Failure      500  {object}  module.HttpErrs
// @Router /api/chum/update [post]
func ApplyUpdate(g *gin.Context) {

	params := new(typeChum.ApplyUpdateParams)
	if g.Bind(params) != nil {
		g.JSON(http.StatusOK, module.ResponseErrorParams("参数不全"))
		return
	}

	token := token2.GetToken(g)

	auth, err := home.HomesMessageGet(token)

	if err != nil {
		g.JSON(http.StatusInternalServerError, module.ResponseServerError("server err"))
		return
	}

	err = chum.NewsApplyUser(params.ID, params.Dispose, auth.Uuid)

	if err != nil {
		g.JSON(http.StatusInternalServerError, module.ResponseServerError("server err new typeUser chums"))
		return
	}

	g.JSON(http.StatusOK, module.ResponseSuccess("添加成功"))
}

// ChumsListGet 获取好友列表
// @Tags Chum
// @Summary 获取好友列表
// @ID ChumsListGet
// @Produce  json
// @Accept  json
// @Param token header  string true "token"
// @Success 200 {object} module.ResponseBodyInString true "JSON数据"
// @Failure      400  {object}  module.HttpErrs
// @Failure      404  {object}  module.HttpErrs
// @Failure      500  {object}  module.HttpErrs
// @Router /api/chum/update [post]
func ChumsListGet() {

}
