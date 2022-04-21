package chum

import (
	"github.com/gin-gonic/gin"
	"modTest/module"
	"net/http"
)

func SearchUser(g *gin.Context) {
	phone := g.Param("phone")

	if phone == "" {
		g.JSON(http.StatusOK, module.ResponseErrorParams("参数不全"))
		return
	}

}
