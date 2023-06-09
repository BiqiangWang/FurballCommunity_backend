// 文件上传模块
package controller

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MultiUpload(c *gin.Context) {
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "文件上传失败： %s", err.Error())
		return
	}
	files := form.File["file"]
	// 保存上传的文件路径
	var returnPaths []string
	for _, file := range files {
		// 生成新的文件名
		ext := filepath.Ext(file.Filename)
		// 通过引入第三方库github.com/google/uuid来生成唯一的UUID字符串
		newFilename := uuid.New().String() + ext
		err = c.SaveUploadedFile(file, "../img/"+newFilename)
		if err != nil {
			c.String(http.StatusBadRequest, "文件上传失败： %s", err.Error())
			return
		}
		// 添加文件路径到切片中
		returnPaths = append(returnPaths, "/img/"+newFilename)
	}
	// 返回文件路径列表
	c.JSON(http.StatusOK, gin.H{
		"paths": returnPaths,
	})
}
