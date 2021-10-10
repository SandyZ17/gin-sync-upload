package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SandyZ17/gin-sync-upload/adapter"
)

func main() {
	router := gin.Default()
	router.POST("/api/file/:filename/:filetype", func(c *gin.Context) {

	})
	router.POST("/api/file/slice/:filename/:filetype", func(c *gin.Context) {
		// Get filename
		fileName := c.Param("filename")
		fileType := c.Param("filetype")
		// 获取body内容
		fileContext := c.Request.Body
		defer fileContext.Close()
		// 处理上传
		err := adapter.UploadSliceFile(fileName, fileType, fileContext)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, "files uploaded!")
	})
	router.Run(":8000")
}
