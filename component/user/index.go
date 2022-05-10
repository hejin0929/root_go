package user

import (
	"modTest/module/user"
	"modTest/service/DB"
)

func SearchUser(uuid string) (*user.User, *user.Message) {

	resUser := new(user.User)

	message := new(user.Message)

	db, _ := DB.CreateDB()

	db.Model(user.User{}).Where("uuid=?", uuid).First(&resUser)

	return nil, nil
}
