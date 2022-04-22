package home

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"modTest/module"
	"modTest/module/center"
	"modTest/service/DB"
	"net/http"
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
