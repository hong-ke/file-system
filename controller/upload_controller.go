package controller

import (
	"filesystem/dig"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary get-hello
// @Tags hello
// @Description
// @Accept json
// @Param Authorization header string true "Header Auth Access Token of User"
// @Param id path string true "群组id"
// @Success  204
// @Router /v1/{hello}  [get]
func GetIndexHTMLController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "文件上传",
	})
}

// 上传文件
func UploadFileController(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	service, err := dig.NewUploadService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	err = service.SaveFile(file, header.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, "upload success !")
}

// 查看文件信息
func GetFileMetaController(ctx *gin.Context) {
	fileSha1 := ctx.Param("file-sha1")
	service, err := dig.NewUploadService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	meta, err := service.GetFileMeta(fileSha1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, meta)
}

// 下载文件
func DownloadFileController(ctx *gin.Context) {
	fileSha1 := ctx.Param("file-sha1")
	service, err := dig.NewUploadService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	data, name, err := service.DownloadFile(fileSha1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Writer.Header().Add("Content-Type", "application/octect-stream")
	ctx.Writer.Header().Add("content-disposition", "attachment; filename=\""+*name+"\"")
	ctx.Data(http.StatusOK, "text/plain; charset=utf-8", data)
}
