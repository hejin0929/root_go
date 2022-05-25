package pathGetCode

import (
	"github.com/gin-gonic/gin"
	"modTest/component/getCode"
	"net/http"
)

// GetPathsCode
// 获取验证码
// @Tags Code
// @Summary 用户验证码登录
// @ID GetPathsCode
// @Produce  json
// @Accept  json
// @Param phone path string true "用户手机号"
// @Success 200 {object} getCode.ResCode true "JSON数据"
// @Router /api/phone_code/typeUser/{phone} [get]
func GetPathsCode(r *gin.Context) {
	phone := r.Query("phone")
	var res getCode.ResCode

	if phone == "" {
		res.MgsCode = 500
		res.MgsText = "手机号码不能为空！"
		res.Body.Code = "0"
		r.JSON(http.StatusOK, res)
		return
	}

	getCode.GetPhoneCodeMessage(r, phone)
}
