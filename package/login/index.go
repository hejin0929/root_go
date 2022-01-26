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
	"regexp"
	"strconv"
	"strings"
	"time"
)

type RepsGetCode struct {
	module.Resp
	// in Body
	Body struct {
		GetPhoneCode
	}
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
	Phone, err := r.Params.Get("name")
	reps := RepsGetCode{}

	myLog := my_log.GetLog()
	defer myLog.CloseFileLog()
	writeLog := myLog.GetLogger()

	if err {
		writeLog.Println(err)
	}

	if RgxPhone(r, Phone) {
		return
	}

	reps.MgsCode = 200

	reps.MgsText = "验证码获取成功"

	var code string

	for i := 0; i < 6; i++ {
		code = code + strconv.Itoa(rand.Intn(10))
	}

	reps.Body.Code = code

	r.JSON(200, reps)

}

// UserBody
// 用户登录成功的返回体
type UserBody struct {
	module.Resp
	Body struct {
		User
	}
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
		writeLog.Println("登录缺少手机号", user.Code)
		r.JSON(500, resp)
		return
	}

	if RgxPhone(r, user.Phone) {
		return
	}

	Users := struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
		Uid      string `json:"uid"`
	}{}

	if user.Code == "" {
		if Users.Password != userP.Password {
			resp.MgsCode = 500
			resp.MgsText = "密码错误。登录失败"
			r.JSON(200, resp)
		}
	}
	claims := &module.JWTClaims{
		UserID:      1,
		Username:    Users.Phone,
		Password:    Users.Password,
		FullName:    Users.Phone,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()

	resp.MgsCode = 200
	resp.MgsText = "登录成功"
	token, _ := token.CreateToken(claims)
	resp.Body.Token = token
	resp.Body.Name = Users.Phone
	resp.Body.Uid = Users.Uid

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
// @Param NewUser body	UserSignType true "JSON数据"
// @Success 200 {object} UserSignReps true "JSON数据"
// @Router /login/user/sign [post]
func SignUser(r *gin.Context) {

	body, err := ioutil.ReadAll(r.Request.Body)

	resp := UserBody{}

	myLog := my_log.GetLog()

	defer myLog.CloseFileLog()

	writeLog := myLog.GetLogger()

	if err != nil {
		writeLog.Println(err)
	}

	newUser := Users{}

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		return
	}

	if RgxPassword(r, newUser.Password) {
		return
	}
	if IsCode(r, newUser.Code) {
		return
	}

	db, err := DB.CreateDB()

	err = db.AutoMigrate(&Users{})
	if err != nil {
		return
	}

	isUser := Users{}

	db.Model(&Users{}).Where("phone=?", newUser.Phone).First(&isUser)

	if isUser.Phone == newUser.Phone {
		resp.MgsCode = 500
		resp.MgsText = "该账号已被注册!"
		r.JSON(500, resp)
		return
	}

	// 建立uuid的地方
	u2 := uuid.NewV4()

	uid := strings.ReplaceAll(u2.String(), "-", "")

	newUser.UUID = uid

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
		//writeLog.Println(isNameErr)
	}

	fmt.Println("this is ??", uid)

	resp.MgsCode = 200
	resp.MgsText = "欢迎你的加入!注册成功!"

	r.JSON(200, resp)
}

//type SetPasswordResp struct {
//	module.Resp
//	Body string `json:"body"`
//}
//
//// SetPasswordUser
//// 设置修改密码的处理函数
//// @Tags Login
//// @Summary 用户休息密码操作
//// @ID SetPasswordUser
//// @Param body body	SetPassword true "JSON数据"
//// @Success 200 {object} SetPasswordResp true "JSON数据"
//// @Router /login/user/set_password [post]
//func SetPasswordUser(r *gin.Context) {
//	body, err := ioutil.ReadAll(r.Request.Body)
//
//	if err != nil {
//		fmt.Println(err)
//		my_log.WriteLog().Println(err)
//	}
//
//	resp := SetPasswordResp{}
//
//	userMse := SetPassword{}
//
//	user := UserName{}
//
//	json.Unmarshal(body, &userMse)
//
//	logs := my_log.GetLog()
//
//	defer logs.CloseFileLog()
//
//	sql := mysql_link.InitDB("root:L6:D*v8GOoif@tcp(127.0.0.1:3306)/user")
//
//	selectNameTemplate := fmt.Sprintf(`SELECT phone, password FROM Users WHERE phone = '%s'`, userMse.Phone)
//
//	sql.SelectOneData(selectNameTemplate).Scan(&user.Phone, &user.Password)
//
//	if user.Phone == "" {
//		resp.MgsCode = 500
//		resp.MgsText = "修改密码异常!"
//		r.JSON(500, resp)
//		logs.GetLogger().Println("数据库未查到用户. ", user.Phone)
//		return
//	}
//
//	if userMse.OldPassword != user.Password {
//		resp.MgsCode = 500
//		resp.MgsText = "密码不正确"
//		resp.Body = "密码不正确修改失败"
//
//		r.JSON(500, resp)
//		return
//	}
//
//	if userMse.NewPassword == user.Password {
//		resp.MgsCode = 500
//		resp.MgsText = "新旧密码不能相同!"
//		resp.Body = "新旧密码相同，请更换其他密码"
//		r.JSON(500, resp)
//		return
//	}
//
//
//	id, err := sql.InsertOneData("update users set password=? where phone=?", logs.GetLogger(), userMse.NewPassword, userMse.Phone)
//
//	if err != nil {
//		logs.GetLogger().Println(err)
//		resp.MgsCode = 500
//		resp.MgsText = "内部错误"
//		resp.Body = err.Error()
//
//		r.JSON(500, resp)
//	}
//
//	resp.MgsCode = 200
//	resp.MgsText = "密码修改成功"
//
//	resp.Body = "请重新进行登录" + strconv.Itoa(int(id))
//
//	r.JSON(200, resp)
//}

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
		r.JSON(500, reps)
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
		r.JSON(500, reps)
		return true
	}

	return false
}

// IsCode
// 验证code验证是否合法
func IsCode(r *gin.Context, code string) bool {
	myLog := my_log.GetLog()
	defer myLog.CloseFileLog()
	writeLog := myLog.GetLogger()
	resp := module.Resp{}

	if code == "" {
		resp.MgsCode = 500
		resp.MgsText = "验证码为空"
		writeLog.Println("登录验证码为空 ---> ", code)
		r.JSON(500, resp)
		return true

	} else if len(code) > 6 {
		resp.MgsCode = 500
		resp.MgsText = "验证码位数错误"
		writeLog.Println("验证码位数错误 ---> ", code)
		r.JSON(500, resp)
		return true
	}

	return false
}

type UserMessage struct {
	Name string `json:"name"`
	Say  int    `json:"say"`
}
