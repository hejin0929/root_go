package home

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"modTest/module"
	"modTest/module/center"
	"modTest/service/DB"
	"net/http"
	"reflect"
	"strings"
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

	resp.Image = "http://localhost:8081/oss" + resp.Image

	g.JSON(http.StatusOK, module.ResponseSuccess(resp))
	return
}

func UpdateUserMessage(g *gin.Context) {

	reqData := new(struct {
		Data UserMessage `json:"data"`
	})
	res := new(MessageUpdate)

	if g.Bind(reqData) != nil {
		res.Res = "参数不全"
		_ = g.Bind(res.Message)
		g.JSON(http.StatusOK, module.ResponseErrorParams(res))
		return
	}

	db, _ := DB.CreateDB()

	user := new(center.Message)

	data := reqData.Data

	user.Uuid = data.Uuid

	db.First(&user)

	types := reflect.TypeOf(&data).Elem()

	values := reflect.ValueOf(&data).Elem()

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

	user.Image = strings.Split(user.Image, "oss")[1]

	res.Res = "修改成功"
	res.Message = reqData.Data

	db.Save(user)

	g.JSON(http.StatusOK, module.ResponseSuccess(res))

}
