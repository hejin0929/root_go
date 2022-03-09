package uploadFiles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"modTest/module/uploadFiles"
	"net/http"
)

// UploadImg
// 上传图片接口
func UploadImg(g *gin.Context) {
	res := uploadFiles.FilesRes{}
	img, err := g.FormFile("img")
	if err != nil {
		res.MgsCode = http.StatusInternalServerError
		res.MgsText = "图片上传失败！！！"
		g.JSON(200, res)
	}
	fmt.Println(img)
}
