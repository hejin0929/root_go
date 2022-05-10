package home

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"modTest/module"
	"modTest/module/user"
	"modTest/service/DB"
	"modTest/types/resHome"
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

	resUser := new(user.Message)

	resp := new(user.Message)

	db.Model(&user.Message{}).First(&resUser, "uuid = ?", id)

	if resUser.Uuid == "" {
		g.JSON(http.StatusOK, module.ResponseErrorParams("查无此用户"))
		return
	}

	data, _ := json.Marshal(resUser)

	_ = json.Unmarshal(data, &resp)

	resp.Image = "http://localhost:8081/oss" + resp.Image

	g.JSON(http.StatusOK, module.ResponseSuccess(resp))
	return
}

func UpdateUserMessage(g *gin.Context) {

	reqData := new(struct {
		Data user.Message `json:"data"`
	})
	res := new(resHome.MessageUpdate)

	if g.Bind(reqData) != nil {
		res.Res = "参数不全"
		_ = g.Bind(res.Message)
		g.JSON(http.StatusOK, module.ResponseErrorParams(res))
		return
	}

	db, _ := DB.CreateDB()

	user := new(user.Message)

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
