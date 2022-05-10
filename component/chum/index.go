package chum

import (
	"encoding/json"
<<<<<<< HEAD
	"github.com/gin-gonic/gin"
	"modTest/module"
	"modTest/module/user"
=======
	"modTest/module/center"
>>>>>>> 7c5cfbafc3b65ab7e1446ab7bf96cee5a60d0051
	"modTest/service/DB"
	chum2 "modTest/types/chum"
	user2 "modTest/types/user"
	"regexp"
)

<<<<<<< HEAD
func SearchUser(g *gin.Context) {
	phone := g.Query("phone")

	if phone == "" {
		g.JSON(http.StatusOK, module.ResponseErrorParams("参数不全"))
		return
	}
=======
func SearchUser(phone string) *user2.Message {
>>>>>>> 7c5cfbafc3b65ab7e1446ab7bf96cee5a60d0051

	db, _ := DB.CreateDB()

	newUser := new(user.User)

<<<<<<< HEAD
	res := new(chum2.ResChum)
=======
	message := new(user2.Message)
>>>>>>> 7c5cfbafc3b65ab7e1446ab7bf96cee5a60d0051

	newUser.Phone = phone

	result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, phone)

<<<<<<< HEAD
		db.Model(&user.User{}).Where("phone=?", phone).First(&newUser)
		res.Body.Source = 1
=======
	if result {
		db.Model(&center.Message{}).Where("phone=?", phone).First(&user)
>>>>>>> 7c5cfbafc3b65ab7e1446ab7bf96cee5a60d0051

	} else {
		db.Model(&user.Message{}).Where("user_id=?", phone).First(&newUser)
	}

	bytes, _ := json.Marshal(newUser)

	_ = json.Unmarshal(bytes, &res.Body.User)

	res.Body.User.Image = "http://localhost:8081/oss" + res.Body.User.Image

<<<<<<< HEAD
	g.JSON(http.StatusOK, module.ResponseSuccess(res.Body))

=======
	return message
>>>>>>> 7c5cfbafc3b65ab7e1446ab7bf96cee5a60d0051
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
