package chum

import (
	"encoding/json"
	"modTest/module/user"
	"modTest/service/DB"
	chum2 "modTest/types/chum"
	user2 "modTest/types/user"
	"regexp"
)

func SearchUser(phone string) *user2.Message {

	db, _ := DB.CreateDB()

	newUser := new(user.User)
	res := new(chum2.ResChum)
	message := new(user2.Message)

	newUser.Phone = phone

	result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, phone)
	if result {
		db.Model(&user.User{}).Where("phone=?", phone).First(&newUser)
		res.Body.Source = 1
	} else {
		db.Model(&user.Message{}).Where("user_id=?", phone).First(&newUser)
	}

	bytes, _ := json.Marshal(newUser)

	_ = json.Unmarshal(bytes, &res.Body.User)

	res.Body.User.Image = "http://localhost:8081/oss" + res.Body.User.Image

	return message
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
