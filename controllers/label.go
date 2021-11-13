package controllers

import (
	"Linkux/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LabelHandler(c *gin.Context) {
	data, err := logic.GetLabelpostList()
	if err != nil {
		zap.L().Error("logic.GetLabelList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func LabelDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetLabelDetail(id)
	if err != nil {
		zap.L().Error("logic.GetLabelDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
