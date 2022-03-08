package login

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"modTest/module"
	"modTest/service/DB"
	"modTest/utlis/my_log"
	"modTest/utlis/token"
	"time"
)

// LoginsUserPassword
// 用户账号密码登录
// @Tags Login
// @Summary 用户登录操作
// @ID LoginsUserPassword
// @Param data body	UserName true "JSON数据"
// @Success 200 {array} UserBody true "JSON数据"
// @Router /api/login/user/login [post]
func LoginsUserPassword(r *gin.Context) {

}

// LoginsUserCode
// 用户手机验证码登录
// @Tags Login
// @Summary 用户登录操作
// @ID LoginsUserCode
// @Param data body	UserCode true "JSON数据"
// @Success 200 {array} UserBody true "JSON数据"
// @Router /api/login/user/login_code [post]
func LoginsUserCode(r *gin.Context) {
	body, err := ioutil.ReadAll(r.Request.Body)
	myLog := my_log.GetLog()
	defer myLog.CloseFileLog()
	writeLog := myLog.GetLogger()

	if err != nil {
		writeLog.Println(err)
		r.JSON(500, "服务器内部错误")
		return
	}

	resp := UserBody{}

	user := UserCode{}

	userP := UserName{}

	userMessage := Users{}

	err = json.Unmarshal(body, &user)

	if err != nil {
		return
	}

	if user.Code != "" {
		if IsCode(r, user.Code) {
			return
		}
	} else {
		json.Unmarshal(body, &userP)
	}

	if user.Phone == "" {
		resp.MgsCode = 500
		resp.MgsText = "手机号码为空"
		//writeLog.Println("登录缺少手机号", user.Code)
		r.JSON(200, resp)
		return
	}

	if RgxPhone(r, user.Phone) {
		return
	}

	db, _ := DB.CreateDB()

	db.Model(Users{}).Where("phone=?", user.Phone).First(&userMessage)

	if userMessage.ID == 0 {
		resp.MgsCode = 500
		resp.MgsText = "该手机好吗未注册"
		r.JSON(200, resp)
		return
	}

	if user.Code == "" {
		if userMessage.Password != userP.Password {
			resp.MgsCode = 500
			resp.MgsText = "密码错误。登录失败"
			r.JSON(200, resp)
		}
	}

	claims := &module.JWTClaims{
		UserID:      1,
		Username:    userMessage.Phone,
		Password:    userMessage.Password,
		FullName:    userMessage.Phone,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()

	resp.MgsCode = 200
	resp.MgsText = "登录成功"

	token, _ := token.CreateToken(claims)

	loginUser := module.ActiveUserLogin{}

	loginUser.Login = true

	loginUser.Phone = user.Phone

	loginUser.Token = token

	login := module.ActiveUserLogin{}

	//db.AutoMigrate(module.ActiveUserLogin{})

	_ = db.Model(&module.ActiveUserLogin{}).Where("phone=?", userMessage.Phone).First(&login).Error

	if login.Phone == "" {
		db.Model(&module.ActiveUserLogin{}).Create(&loginUser)
		fmt.Println("this is a Error ?? ")
	} else {
		db.Model(&module.ActiveUserLogin{}).Where("phone=?", userMessage.Phone).Updates(map[string]interface{}{"login": loginUser.Login, "token": loginUser.Token, "updated_at": time.Now()})
	}

	resp.Body.Token = token
	resp.Body.Name = userMessage.Phone
	resp.Body.Uid = userMessage.UUID

	r.JSON(200, resp)
}
