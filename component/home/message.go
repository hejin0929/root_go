package home

import (
	"encoding/json"
	"fmt"
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

	resUser := new(user.MessageSql)

	resp := new(user.Message)

	db.Model(&user.MessageSql{}).First(&resUser, "uuid = ?", id)

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

	reqData := new(user.Message)
	res := new(resHome.MessageUpdate)

	if g.Bind(&reqData) != nil {
		res.Res = "参数不全"
		_ = g.Bind(res.Message)
		g.JSON(http.StatusOK, module.ResponseErrorParams(res))
		return
	}

	db, _ := DB.CreateDB()

	user := new(user.MessageSql)

	user.Uuid = reqData.Uuid

	db.First(&user)

	types := reflect.TypeOf(&reqData).Elem().Elem()

	values := reflect.ValueOf(&reqData).Elem().Elem()

	userT := reflect.TypeOf(&user.Message).Elem()

	userV := reflect.ValueOf(&user.Message).Elem()

	for i := 0; i < values.NumField(); i++ {
		//fmt.Println("this is a ?? ", types.Field(i).Name, i)
		for l := 0; l < userT.NumField(); l++ {
			fmt.Println("this is a ?? ", userT.Field(l).Name, l)
			if types.Field(i).Name == userT.Field(l).Name {

				if values.Field(i).String() != "" || values.Field(i).String() != "nil" || types.Field(i).Name != "uuid" {
					userV.Field(l).Set(values.Field(i))
				}

			}
		}
	}

	user.Image = strings.Split(user.Image, "oss")[1]

	res.Res = "修改成功"

	res.Message = *reqData

	db.Save(user)

	g.JSON(http.StatusOK, module.ResponseSuccess(res))

}
