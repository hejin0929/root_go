package user

import (
	"encoding/json"
	"modTest/module/user"
	"modTest/service/DB"
)

func SearchUser(uuid string) (*user.MessageSql, *user.Message) {

	resUser := new(user.MessageSql)

	message := new(user.Message)

	db, _ := DB.CreateDB()

	db.Model(user.MessageSql{}).Where("uuid=?", uuid).First(resUser)

	bytes, _ := json.Marshal(resUser)

	_ = json.Unmarshal(bytes, &message)

	return resUser, message
}
