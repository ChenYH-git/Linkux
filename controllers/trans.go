package controllers

import (
	"Linkux/logic"
	"Linkux/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateTransTaskHandler(c *gin.Context) {
	p := new(models.Trans)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON failed", zap.Any("Err", err))
		zap.L().Error("create trans failed")
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := logic.CreateTrans(p); err != nil {
		zap.L().Error("logic.CreateTrans(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, gin.H{
		"msg": "create trans-help success",
	})
}

func GetTransTaskHandler(c *gin.Context) {
	data, err := logic.GetTransTask()
	if err != nil {
		zap.L().Error("logic.GetTransTask() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

func GetTransExistHandler(c *gin.Context) {
	p := &models.ParamPostList{
		Page: 1,
		Size: 10,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetTransExistHandler get query err: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetTransExist(p)
	if err != nil {
		zap.L().Error("logic.GetTransExist(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}
