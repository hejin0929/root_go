package getCode

import (
	"github.com/gin-gonic/gin"
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
// @Success 200 {object} ResCode true "JSON数据"
// @Router /code/user/{phone} [get]
func GetPathsCode(r *gin.Context) {
	phone := r.Query("phone")
	if phone == "" {
		r.JSON(http.StatusOK, "1")
		return
	}
}
