package chum

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"modTest/module"
	"modTest/module/user"
	"modTest/service/DB"
	chum2 "modTest/types/chum"
	"net/http"
	"regexp"
)

func SearchUser(g *gin.Context) {
	phone := g.Query("phone")

	if phone == "" {
		g.JSON(http.StatusOK, module.ResponseErrorParams("参数不全"))
		return
	}

	db, _ := DB.CreateDB()

	newUser := new(user.User)

	res := new(chum2.ResChum)

	newUser.Phone = phone

	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)

	if rgx.MatchString(phone) {

		db.Model(&user.User{}).Where("phone=?", phone).First(&newUser)
		res.Body.Source = 1

	} else {
		db.Model(&user.Message{}).Where("user_id=?", phone).First(&newUser)
	}

	bytes, _ := json.Marshal(newUser)

	_ = json.Unmarshal(bytes, &res.Body.User)

	res.Body.User.Image = "http://localhost:8081/oss" + res.Body.User.Image

	g.JSON(http.StatusOK, module.ResponseSuccess(res.Body))

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
