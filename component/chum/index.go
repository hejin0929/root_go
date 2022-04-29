package chum

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"modTest/component/home"
	"modTest/module"
	"modTest/module/center"
	"modTest/service/DB"
	chum2 "modTest/types/chum"
	"net/http"
	"regexp"
)

func SearchUser(g *gin.Context) {
	phone := g.Param("phone")

	if phone == "" {
		g.JSON(http.StatusOK, module.ResponseErrorParams("参数不全"))
		return
	}

	db, _ := DB.CreateDB()

	user := new(center.Message)

	message := new(home.UserMessage)

	user.Phone = phone

	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)

	if rgx.MatchString(phone) {

		db.Model(&center.Message{}).Where("phone=?", phone).First(&user)
	} else {
		db.Model(&center.Message{}).Where("user_id=?", phone).First(&user)
	}

	bytes, _ := json.Marshal(user)

	_ = json.Unmarshal(bytes, &message)

	g.JSON(http.StatusOK, module.ResponseSuccess(message))

}

// AddChumUser 函数不在返回请求
func AddChumUser(params *chum2.AddReq) error {

	newChum := new(chum2.UserChum)

	bytes, _ := json.Marshal(params)

	_ = json.Unmarshal(bytes, &newChum)

	db, _ := DB.CreateDB()

	_ = db.AutoMigrate(chum2.UserChum{})

	db.Create(&newChum)

	return nil
}
