package login

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"modTest/module"
	login2 "modTest/module/login"
	"modTest/service/DB"
	"modTest/service/redis"
	"modTest/utlis/my_log"
	"modTest/utlis/token"
	"net/http"
	"time"
)

// LoginsUserPassword
// 密码登陆处理逻辑
func LoginsUserPassword(g *gin.Context) {

	//user := new(login2.UserName)
	//if g.Bind(user) != nil {
	//	g.JSON(http.StatusBadRequest, "参数错误")
	//}

	bytes, err := ioutil.ReadAll(g.Request.Body)

	reqData := struct {
		Data login2.UserName
	}{}

	_ = json.Unmarshal(bytes, &reqData)

	user := reqData.Data

	fmt.Println("this is value ?? ", string(bytes))

	res := struct {
		module.Resp
		Body login2.User `json:"body"`
	}{}

	mqlUser := Users{}

	db, err := DB.CreateDB()

	if err != nil {
		res.MgsCode = 500
		res.MgsText = err.Error()

		g.JSON(200, res)
	}

	fmt.Println("this is a ?? ", user)

	err = db.Model(Users{}).Where("phone=?", user.Phone).First(&mqlUser).Error

	if mqlUser.Phone == "" {
		res.MgsCode = 500
		res.MgsText = "暂未注册"
		g.JSON(http.StatusBadRequest, res)
		return
	}

	if mqlUser.Password != user.Password {
		res.MgsCode = 500
		res.MgsText = "密码错误"
		g.JSON(http.StatusOK, res)
		return
	}

	claims := &module.JWTClaims{
		UserID:      mqlUser.ID,
		Username:    user.Phone,
		Password:    user.Password,
		FullName:    user.Phone,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()

	token, _ := token.CreateToken(claims)

	res.Body.Token = token
	res.Body.Uid = mqlUser.UUID
	res.Body.Name = mqlUser.Phone

	redisToken := redis.CreateRedis(0)

	redisToken.Set(context.Background(), mqlUser.UUID, token, 0)

	res.MgsCode = 200
	res.MgsText = "登陆成功"

	g.JSON(http.StatusOK, res)

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

	user := login2.UserCode{}

	userP := login2.UserName{}

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
		UserID:      userMessage.ID,
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
	} else {
		db.Model(&module.ActiveUserLogin{}).Where("phone=?", userMessage.Phone).Updates(map[string]interface{}{"login": loginUser.Login, "token": loginUser.Token, "updated_at": time.Now()})
	}

	resp.Body.Token = token
	resp.Body.Name = userMessage.Phone
	resp.Body.Uid = userMessage.UUID

	r.JSON(200, resp)
	return
}
