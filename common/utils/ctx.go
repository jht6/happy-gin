package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CtxResOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    data,
	})
}

func CtxResErr(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": msg,
	})
}

func CtxResErrWithCode(ctx *gin.Context, msg string, code int) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
	})
}
