package uploadFiles

import (
	"github.com/gin-gonic/gin"
	"modTest/component/uploadFiles"
)

// UploadImages
// ** 用户上传图片逻辑
// @Tags Upload
// @Summary 上传图片操作
// @ID UploadImages
// @Param image formData file true "file"
// @Success 200 {object} uploadFiles.FilesRes true "JSON数据"
// @Router /api/upload/images [post]
func UploadImages(g *gin.Context) {
	uploadFiles.UploadImg(g)
}
