package handler

import (
	"happy-gin/common/utils"

	"github.com/gin-gonic/gin"
)

func DemoOk(ctx *gin.Context) {
	utils.CtxResOk(ctx, gin.H{
		"num": 6,
	})
}

func DemoErr(ctx *gin.Context) {
	utils.CtxResErr(ctx, "demo err")
}

func DemoErrWithCode(ctx *gin.Context) {
	utils.CtxResErrWithCode(ctx, "err with code", 6)
}
