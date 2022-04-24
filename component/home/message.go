package home

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"modTest/module"
	"modTest/module/center"
	"modTest/service/DB"
	"net/http"
	"reflect"
)

func GetUserMessage(g *gin.Context) {
	id := g.Query("id")

	if id == "" {
		g.JSON(http.StatusOK, module.ResponseSuccess("ID 缺少"))
		return
	}

	db, _ := DB.CreateDB()

	user := new(center.Message)

	resp := new(UserMessage)

	db.Model(&center.Message{}).First(&user, "uuid = ?", id)

	if user.Uuid == "" {

		g.JSON(http.StatusOK, module.ResponseErrorParams("查无此用户"))
		return
	}

	data, _ := json.Marshal(user)

	_ = json.Unmarshal(data, &resp)

	resp.Image = "http://localhost:8081/oss/" + resp.Image

	g.JSON(http.StatusOK, module.ResponseSuccess(resp))
	return
}

func UpdateUserMessage(g *gin.Context) {

	reqData := new(UserMessage)
	res := new(MessageUpdateRes)

	if g.Bind(reqData) != nil {
		res.MgsText = "Error"
		res.MgsCode = http.StatusPreconditionFailed
		res.Body.Res = "参数不全"
		_ = g.Bind(res.Body.Message)
		g.JSON(http.StatusOK, res)
		return
	}

	db, _ := DB.CreateDB()

	user := new(center.Message)

	user.Uuid = reqData.Uuid

	db.First(&user)

	types := reflect.TypeOf(reqData).Elem()

	values := reflect.ValueOf(reqData).Elem()

	userT := reflect.TypeOf(user).Elem()

	userV := reflect.ValueOf(user).Elem()

	for i := 0; i < values.NumField(); i++ {
		for l := 0; l < userT.NumField(); l++ {
			if types.Field(i).Name == userT.Field(l).Name {

				if values.Field(i).String() != "" || values.Field(i).String() != "nil" || types.Field(i).Name != "uuid" {
					userV.Field(l).Set(values.Field(i))
				}

			}
		}
	}

	db.Save(&userV)

}
