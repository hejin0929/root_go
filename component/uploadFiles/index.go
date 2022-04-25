package uploadFiles

import (
	"github.com/gin-gonic/gin"
	"modTest/module/uploadFiles"
	"net/http"
	"os"
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

	var url = path + "/static/images/" + filename

	err = g.SaveUploadedFile(file, url)
	if err != nil {
		res.MgsCode = http.StatusInternalServerError
		res.MgsText = "服务端保存图片错误！！！"
		g.JSON(http.StatusOK, res)
		return
	}

	res.MgsCode = http.StatusOK
	res.MgsText = "图片上传成功"
	res.Body = "http://localhost:8081/oss/images/" + filename

	g.JSON(http.StatusOK, res)
	return
}

// DeleteImgDelete
// 删除照片逻辑
func DeleteImgDelete(g *gin.Context) {
	name := g.Query("name")

	res := uploadFiles.FileImageDelete{}

	if name == "" {
		res.MgsCode = http.StatusInternalServerError
		res.MgsText = "删除图片名称为空！！！"
		g.JSON(http.StatusOK, res)
		return
	}

	path, _ := filepath.Abs("./")

	err := os.Remove(path + "/static/images/" + name)
	if err != nil {
		res.MgsCode = http.StatusInternalServerError
		res.MgsText = "图片删除失败！！！"
		g.JSON(http.StatusOK, res)
		return
	}

	res.MgsCode = http.StatusOK
	res.MgsText = "success"
	res.Body = "删除成功！！！"

	g.JSON(http.StatusOK, res)
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

// UploadVideoDelete
// 删除视频接口处理
func UploadVideoDelete(g *gin.Context) {
	name := g.Query("name")

	res := uploadFiles.FileImageDelete{}

	if name == "" {
		res.MgsCode = http.StatusLocked
		res.MgsText = "删除视频不能为空！！！"
		g.JSON(http.StatusOK, res)

		return
	}

	path, _ := filepath.Abs("./")

	err := os.Remove(path + "/static/videos/" + name)

	if err != nil {

		res.MgsCode = http.StatusInternalServerError

		res.MgsText = "删除视频失败！！！"

		g.JSON(http.StatusOK, res)

		return
	}

	res.MgsCode = http.StatusOK

	res.MgsText = "success"

	res.Body = "删除视频成功！！！"

	g.JSON(http.StatusOK, res)

}
