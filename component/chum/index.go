package chum

import (
	"encoding/json"
	"modTest/module/center"
	"modTest/service/DB"
	chum2 "modTest/types/chum"
	user2 "modTest/types/user"
	"regexp"
)

func SearchUser(phone string) *user2.Message {

	db, _ := DB.CreateDB()

	user := new(center.Message)

	message := new(user2.Message)

	user.Phone = phone

	result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, phone)

	if result {
		db.Model(&center.Message{}).Where("phone=?", phone).First(&user)

	} else {
		db.Model(&center.Message{}).Where("user_id=?", phone).First(&user)
	}

	bytes, _ := json.Marshal(user)

	_ = json.Unmarshal(bytes, &message)

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
