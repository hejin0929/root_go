package login

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"math/rand"
	"modTest/module"
	"modTest/service/DB"
	"modTest/utlis/my_log"
	"modTest/utlis/token"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type RepsGetCode struct {
	module.Resp
	// in Body
	Body GetPhoneCode `json:"body"` // 返回体
}

// GetSingCode
// 获取登录验证码
// GetSingCode
// @Tags Login
// @Summary 用户验证码登录
// @ID GetSingCode
// @Produce  json
// @Accept  json
// @Param phone path string true "用户手机号"
// @Success 200 {object} RepsGetCode true "JSON数据"
// @Router /login/user/{phone} [get]
func GetSingCode(r *gin.Context) {
	Phone := r.Query("phone")
	reps := RepsGetCode{}

	myLog := my_log.GetLog()
	defer myLog.CloseFileLog()
	//writeLog := myLog.GetLogger()

	//if err {
	//	writeLog.Println(err)
	//}

	if RgxPhone(r, Phone) {
		return
	}

	reps.MgsCode = 200

	reps.MgsText = "验证码获取成功"

	var code string

	for i := 0; i < 6; i++ {
		code = code + strconv.Itoa(rand.Intn(10))
	}

	db, errDB := DB.CreateDB()

	if errDB != nil {
		r.JSON(http.StatusInternalServerError, errDB)
		return
	}

	crateDBErr := db.AutoMigrate(&UserCode{})

	if crateDBErr != nil {
		r.JSON(http.StatusInternalServerError, crateDBErr)
		return
	}

	isCode := UserCode{}

	_ = db.Model(&UserCode{}).Where("phone=?", Phone).First(&isCode).Error

	reps.Body.Code = code

	if isCode.Code != "" {
		db.Model(&UserCode{}).Where("phone=?", Phone).Update("code", code)

		r.JSON(200, reps)
		return
	}

	db.Model(&UserCode{}).Create(UserCode{Phone: Phone, Code: code})

}

// UserBody
// 用户登录成功的返回体
type UserBody struct {
	module.Resp
	Body User
}

// LoginsUser
// 用户账号密码登录
// @Tags Login
// @Summary 用户登录操作
// @ID LoginsUser
// @Param data body	UserCode true "JSON数据"
// @Success 200 {array} UserBody true "JSON数据"
// @Router /login/user/login [post]
func LoginsUser(r *gin.Context) {
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

	db.AutoMigrate(module.ActiveUserLogin{})

	_ = db.Model(&module.ActiveUserLogin{}).Where("phone=?", userMessage.Phone).First(&login).Error

	if login.Phone == "" {
		db.Model(&module.ActiveUserLogin{}).Create(loginUser)
	} else {
		db.Model(&userMessage).Where("phone=?", userMessage.Phone).Updates(loginUser)
	}

	resp.Body.Token = token
	resp.Body.Name = userMessage.Phone
	resp.Body.Uid = userMessage.UUID

	r.JSON(200, resp)
}

type UserSignType struct {
	UserName
	// 手机验证码
	Code string `json:"code" binding:"required"`
}

type UserSignReps struct {
	module.Resp
	Body string `json:"body"`
}

type Users struct {
	module.Model
	UserName
	UUID string `json:"uuid"` // uuid
	Code string `json:"code"` // Code
}

// SignUser
// 注册用户处理逻辑
// @Tags Login
// @Summary 用户登录操作
// @ID SignUser
// @Param data body	UserSignType true "JSON数据"
// @Success 200 {object} UserSignReps true "JSON数据"
// @Router /login/user/sign [post]
func SignUser(r *gin.Context) {

	body, err := ioutil.ReadAll(r.Request.Body)
	//fmt.Println("this is body ?? value", string(body))

	resp := UserBody{}

	myLog := my_log.GetLog()

	defer myLog.CloseFileLog()

	writeLog := myLog.GetLogger()

	if err != nil {
		writeLog.Println(err)
	}

	data := struct {
		Data UserSignType `json:"data"`
	}{}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return
	}

	newUser := Users{}

	if RgxPassword(r, data.Data.Password) {
		return
	}

	if IsCode(r, data.Data.Code) {
		return
	}

	if RgxPhone(r, data.Data.Phone) {
		return
	}

	db, err := DB.CreateDB()

	db.Model(Users{})

	codeRgx := UserCode{}

	db.Model(&UserCode{}).Where("phone=?", data.Data.Phone).First(&codeRgx)

	if codeRgx.Code != data.Data.Code {
		resp.MgsCode = 500
		resp.MgsText = "验证码错误!"
		r.JSON(200, resp)
		return
	}

	err = db.AutoMigrate(&Users{})

	isUser := Users{}

	_ = db.Model(&Users{}).Where("phone=?", data.Data.Phone).First(&isUser).Error

	if err != nil {

		r.JSON(500, "服务器内部错误")
		return
	}

	if isUser.Phone == data.Data.Phone {
		resp.MgsCode = 500
		resp.MgsText = "该账号已被注册!"
		r.JSON(200, resp)
		return
	}

	// 建立uuid的地方
	u2 := uuid.NewV4()

	uid := strings.ReplaceAll(u2.String(), "-", "")

	newUser.UUID = uid
	newUser.Phone = data.Data.Phone
	newUser.Code = data.Data.Code
	newUser.Password = data.Data.Password

	db.Create(&newUser)

	if err != nil {
		r.JSON(500, err)
		return
	}

	if err != nil {
		return
	}

	if err != nil {
		fmt.Println("ERR IS ?? ", err)
		return
	}

	if err != nil {
		writeLog.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	resp.MgsCode = 200
	resp.MgsText = "欢迎你的加入!注册成功!"

	r.JSON(200, resp)
}

var (
	Secret     = "dong_tech" // 加盐
	ExpireTime = 3600        // token有效期
)

// RgxPhone
// 检查手机验证码是否合法化
func RgxPhone(r *gin.Context, phone string) bool {
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)

	reps := module.Resp{}

	myLog := my_log.GetLog()
	defer myLog.CloseFileLog()
	writeLog := myLog.GetLogger()

	if !rgx.MatchString(phone) {
		writeLog.Println("手机号码不合法", "用户 -->", phone)
		reps.MgsCode = 500
		reps.MgsText = "手机号码不合法"
		r.JSON(200, reps)
		return true
	}

	return false
}

// RgxPassword
// 检查密码是否合法化
func RgxPassword(r *gin.Context, password string) bool {
	passwordReg := "/^(![0-9]+$)(![a-zA-Z]+$)[0-9A-Za-z]{6,10}$/"
	rgx := regexp.MustCompile(passwordReg)
	myLog := my_log.GetLog()
	defer myLog.CloseFileLog()
	writeLog := myLog.GetLogger()

	reps := module.Resp{}

	if rgx.MatchString(password) {
		writeLog.Println("密码不合法", password)
		reps.MgsCode = 500
		reps.MgsText = "密码不合法"
		r.JSON(200, reps)
		return true
	}

	return false
}

// IsCode
// 验证code验证是否合法
func IsCode(r *gin.Context, code string) bool {
	myLog := my_log.GetLog()
	defer myLog.CloseFileLog()
	resp := module.Resp{}

	if code == "" {
		resp.MgsCode = 500
		resp.MgsText = "验证码为空"
		r.JSON(200, resp)
		return true

	} else if len(code) != 6 {
		resp.MgsCode = 500
		resp.MgsText = "验证码错误"
		r.JSON(500, resp)
		return true
	}

	return false
}

type UserMessage struct {
	Name string `json:"name"`
	Say  int    `json:"say"`
}
