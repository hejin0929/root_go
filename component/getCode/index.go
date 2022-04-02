package getCode

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"modTest/component/login"
	login2 "modTest/module/login"
	"modTest/service/DB"
	"modTest/utlis/my_log"
	"net/http"
	"strconv"
)

// GetPhoneCodeMessage 获取手机验证码
func GetPhoneCodeMessage(r *gin.Context, phone string) {

	reps := ResCode{}

	myLog := my_log.GetLog()

	defer myLog.CloseFileLog()

	if login.RgxPhone(r, phone) {
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

	crateDBErr := db.AutoMigrate(&login2.UserCode{})

	if crateDBErr != nil {
		r.JSON(http.StatusInternalServerError, crateDBErr)
		return
	}

	isCode := UserCode{}

	db.Model(&UserCode{}).Where("phone=?", phone).First(&isCode)

	//if err != nil {
	//	reps.MgsCode = 500
	//	reps.MgsText = err.Error()
	//	r.JSON(200, reps)
	//	return
	//}

	reps.Body.Code = code

	if isCode.Code != "" {
		db.Model(&UserCode{}).Where("phone=?", phone).Update("code", code)
		r.JSON(200, reps)
		db.Model(&UserCode{}).Create(UserCode{Phone: phone, Code: code})
		return
	}

	db.Model(&UserCode{}).Create(UserCode{Phone: phone, Code: code})
	r.JSON(200, reps)

}
