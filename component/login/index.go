package login

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"math/rand"
	"modTest/module"
	"modTest/module/login"
	"modTest/module/user"
	"modTest/service/DB"
	"modTest/utlis/my_log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type RepsGetCode struct {
	module.Resp
	// in Body
	Body login.GetPhoneCode `json:"body"` // 返回体
}

// GetSingCode
// 新增用户验证码
func GetSingCode(r *gin.Context) {
	Phone := r.Query("phone")

	reps := RepsGetCode{}

	myLog := my_log.GetLog()

	defer myLog.CloseFileLog()

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

	crateDBErr := db.AutoMigrate(&login.UserCode{})

	if crateDBErr != nil {
		r.JSON(http.StatusInternalServerError, crateDBErr)
		return
	}

	isCode := login.UserCode{}

	_ = db.Model(&login.UserCode{}).Where("phone=?", Phone).First(&isCode).Error

	reps.Body.Code = code

	if isCode.Code != "" {
		db.Model(&login.UserCode{}).Where("phone=?", Phone).Update("code", code)

		r.JSON(200, reps)
		return
	}

	db.Model(&login.UserCode{}).Create(login.UserCode{Phone: Phone, Code: code})

}

// UserBody
// 用户登录成功的返回体
type UserBody struct {
	module.Resp
	Body login.User `json:"body"`
}

type UserSignType struct {
	login.UserName
	// 手机验证码
	Code string `json:"code" binding:"required"`
}

type UserSignReps struct {
	module.Resp
	Body string `json:"body"`
}

type Users struct {
	module.Model
	login.UserName
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
// @Router /api/pathLogin/typeUser/sign [post]
func SignUser(r *gin.Context) {

	body, err := ioutil.ReadAll(r.Request.Body)

	resp := UserBody{}

	myLog := my_log.GetLog()

	defer myLog.CloseFileLog()

	writeLog := myLog.GetLogger()

	if err != nil {
		writeLog.Println(err)
	}

	inData := UserSignType{}

	err = json.Unmarshal(body, &inData)

	if err != nil {
		return
	}

	newUser := Users{}

	if RgxPassword(r, inData.Password) {
		return
	}

	if IsCode(r, inData.Code) {
		return
	}

	if RgxPhone(r, inData.Phone) {
		return
	}

	db, err := DB.CreateDB()

	db.Model(Users{})

	codeRgx := login.UserCode{}

	db.Model(&login.UserCode{}).Where("phone=?", inData.Phone).First(&codeRgx)

	if codeRgx.Code != inData.Code {
		resp.MgsCode = 500
		resp.MgsText = "验证码错误!"
		r.JSON(200, resp)
		return
	}

	err = db.AutoMigrate(&Users{})

	isUser := Users{}

	_ = db.Model(&Users{}).Where("phone=?", inData.Phone).First(&isUser).Error

	if err != nil {

		r.JSON(500, "服务器内部错误")
		return
	}

	if isUser.Phone == inData.Phone {
		resp.MgsCode = 500
		resp.MgsText = "该账号已被注册!"
		r.JSON(200, resp)
		return
	}

	// 建立uuid的地方
	u2 := uuid.NewV4()

	uid := strings.ReplaceAll(u2.String(), "-", "")

	newUser.UUID = uid
	newUser.Phone = inData.Phone
	newUser.Code = inData.Code
	newUser.Password = inData.Password

	db.Create(&newUser)

	if err != nil {
		r.JSON(500, err)
		return
	}

	if err != nil {
		resp.MgsCode = 500
		resp.MgsText = err.Error()
		r.JSON(200, resp)
		return
	}

	message := new(user.MessageSql)

	message.Uuid = newUser.UUID
	message.Phone = newUser.Phone
	message.Image = "/images/001.jpg"
	message.Name = "user_" + newUser.Phone
	message.UserID = "id_" + newUser.Phone
	message.Sex = 2

	db.AutoMigrate(user.MessageSql{})

	db.Model(&user.MessageSql{}).Create(&message)

	resp.MgsCode = 200
	resp.MgsText = "欢迎你的加入!注册成功!"

	r.JSON(200, resp)
}

var (
	Secret     = "dong_tech" // 加盐
	ExpireTime = 36000       // token有效期
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
