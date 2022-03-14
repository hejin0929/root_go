package uploadFiles

import (
	"github.com/gin-gonic/gin"
	"modTest/module/uploadFiles"
	"net/http"
	"path/filepath"
)

// UploadImg
// 上传图片接口
func UploadImg(g *gin.Context) {
	res := uploadFiles.FilesRes{}
	file, err := g.FormFile("image")
	if err != nil {
		res.MgsCode = http.StatusInternalServerError
		res.MgsText = "图片上传失败！！！"
		g.JSON(200, res)
		return
	}

	//fmt.Println("this is a", file)

	filename := file.Filename
	path, _ := filepath.Abs("./")
	if err != nil {
		res.MgsCode = http.StatusInternalServerError
		res.MgsText = "内部构建错误！！！"
		g.JSON(200, res)
		return
	}

	var url = path + "/static/" + filename

	err = g.SaveUploadedFile(file, url)
	if err != nil {
		res.MgsCode = http.StatusInternalServerError
		res.MgsText = "服务端保存图片错误！！！"
		g.JSON(http.StatusOK, res)
		return
	}

	res.MgsCode = http.StatusOK
	res.MgsText = "图片上传成功"
	res.Body = "http://localhost:8081/oss/" + filename

	g.JSON(http.StatusOK, res)
	return
}

// DeleteImg
// 删除照片逻辑
func DeleteImg() {

}

// UploadVideo
// 视频上传接口处理
func UploadVideo(g *gin.Context) {
	res := uploadFiles.FilesRes{}

	video, err := g.FormFile("videos")

	if err != nil {
		res.MgsCode = http.StatusInternalServerError
		res.MgsText = "视频处理失败！！！"
		g.JSON(http.StatusOK, res)
		return
	}

	videoName := video.Filename

	path, _ := filepath.Abs("./")

	url := path + "/static/videos/" + videoName

	err = g.SaveUploadedFile(video, url)

	if err != nil {
		res.MgsCode = http.StatusInternalServerError
		res.MgsText = "视频储存出错！！！"
		g.JSON(http.StatusOK, res)
		return
	}

	res.MgsCode = http.StatusOK

	res.MgsText = "success"

	res.Body = "http://localhost:8081/oss/videos/" + videoName

	g.JSON(http.StatusOK, res)

	return
}