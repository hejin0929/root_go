package chum

import (
	"encoding/json"
	"errors"
	"fmt"
	user3 "modTest/component/user"
	"modTest/module/chum"
	"modTest/module/user"
	"modTest/service/DB"
	chum2 "modTest/types/chum"
	"regexp"
	"time"
)

func SearchUser(phone string) *chum2.ResChum {

	db, _ := DB.CreateDB()

	newUser := new(user.MessageSql)

	res := new(chum2.ResChum)

	res.Body.Relation = chum2.NotRelation

	newUser.Phone = phone

	result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, phone)

	if result {
		db.Model(&user.MessageSql{}).Where("phone=?", phone).First(&newUser)
		res.Body.Source = 1
	} else {
		db.Model(&user.MessageSql{}).Where("user_id=?", phone).First(&newUser)
	}

	bytes, _ := json.Marshal(newUser)

	_ = json.Unmarshal(bytes, &res.Body.User)

	var relation chum.UserChumSql

	db.Model(&chum.UserChumSql{}).First(&relation, "user_id=?", res.Body.User.UserID)

	if relation.ID != 0 {
		res.Body.Relation = chum2.Relation
	}

	res.Body.User.Image = "http://localhost:8081/oss" + res.Body.User.Image

	return res
}

// AddChumUser 函数不在返回请求
func AddChumUser(params *chum2.AddReq) error {

	newChum := new(chum.ApplyChumSql)

	bytes, _ := json.Marshal(params)

	_ = json.Unmarshal(bytes, &newChum)

	_, UserMessage := user3.SearchUser(params.Uuid)

	newChum.Uuid = UserMessage.UserID

	newChum.Hello = params.Uuid

	db, _ := DB.CreateDB()

	_ = db.AutoMigrate(chum.ApplyChumSql{})

	times := time.Now()

	now := times.Format("2006-01-02 15:04:05")

	start := time.Date(times.Year(), times.Month(), times.Day(), 0, 0, 0, 0, times.Location()).Format("2006-01-02 15:04:05")

	var applyList []chum.ApplyChumSql

	oldTime := time.Date(times.Year(), times.Month(), times.Day()-3, 0, 0, 0, 0, times.Location()).Format("2006-01-02 15:04:05")

	db.Unscoped().Where("created_at < ?", oldTime).Delete(&chum.ApplyChumSql{})

	db.Model(chum.ApplyChumSql{}).Where("created_at BETWEEN ? AND ?", start, now).Find(&applyList)

	if len(applyList) < 5 {
		db.Create(&newChum)
	} else {
		return errors.New("添加次数已限制")
	}

	return nil
}

func GetApplyUser(uuid string) {
	db, _ := DB.CreateDB()

	var applyList []chum.ApplyChumSql

	db.Model(chum.ApplyChumSql{}).Where("friend_id = ?", uuid).Find(&applyList)

	fmt.Println(applyList)

}

func UpdateApplyUser(userId string) { // 更新好友审核

}
