package webserver

import (
	"github.com/gin-gonic/gin"
	"hsplugin/biz"
	"net/http"
)

// 文件处理器

type FileController struct {
}

// DownloadFile 获取文件内容
func DownloadFile(ctx *gin.Context) {
	fileName := ctx.Param("fileName")
	filePath := ctx.Param("filePath")
	fileContentDisposition := "attachment;filename=\"" + fileName + "\""
	ctx.Header("Content-Type", "plain")
	ctx.Header("Content-Disposition", fileContentDisposition)
	ctx.Data(http.StatusOK, "plain", biz.ReadFromDisk(filePath))
}

func (fileController *FileController) Router(server *WebServer) {
	server.router.GET("/download", DownloadFile)
}
